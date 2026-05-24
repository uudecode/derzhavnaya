package translation

import (
	"Derzhavnaya/internal/config"
	"Derzhavnaya/internal/db"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewTranslationService),
)

type Translator interface {
	Translate(ctx context.Context, text string, targetLang string) (string, error)
}

type LiteLLMClient struct {
	url        string
	key        string
	model      string
	prompts    map[string]string
	client     *http.Client
	glossaryEn map[string]string
	glossaryFr map[string]string
}

func NewTranslationService(lc fx.Lifecycle, cfg *config.Config, queries *db.Queries) (Translator, error) {
	timeout := time.Duration(cfg.Translation.TimeoutSec) * time.Second

	service := &LiteLLMClient{
		url:        cfg.Translation.LiteLLMUrl,
		key:        cfg.Translation.LiteLLMKey,
		client:     &http.Client{Timeout: timeout},
		model:      cfg.Translation.ModelName,
		prompts:    cfg.Translation.Prompts,
		glossaryEn: make(map[string]string),
		glossaryFr: make(map[string]string),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			enRows, err := queries.GetEnglishGlossary(ctx)
			if err != nil {
				return fmt.Errorf("failed to load EN glossary on start: %w", err)
			}

			for _, r := range enRows {
				service.glossaryEn[r.RuTerm] = r.EnTrans.String
			}

			frRows, err := queries.GetFrenchGlossary(ctx)
			if err != nil {
				return fmt.Errorf("failed to load FR glossary on start: %w", err)
			}

			for _, r := range frRows {
				service.glossaryFr[r.RuTerm] = r.FrTrans.String
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return service, nil
}

func (s *LiteLLMClient) Translate(ctx context.Context, text string, targetLang string) (string, error) {
	whitelist := map[string]bool{
		"беда":            true,
		"батюшка":         true,
		"болящий":         true,
		"записочка":       true,
		"сорокоуст":       true,
		"бабка":           true,
		"руками разводят": true,
		"виновата перед Господом":  true,
		"о здравии":                true,
		"по благословению батюшки": true,
	}
	var glossaryHeader string
	var glossaryRules string
	var glossary map[string]string
	if targetLang == "en" {
		glossaryRules = "MANDATORY GLOSSARY RULES:\n- Use glossary entries only when the Russian term or phrase appears in the source text.\n- If an entry contains \"/\", choose exactly one option.\n- Never copy glossary alternatives such as \"A / B\" into the translation.\n- Preserve the original subject and object of each sentence.\n"
		glossaryHeader = `MANDATORY GLOSSARY: When these Russian expressions occur in the source text, use the corresponding translation. Do not reinterpret them. Do not turn statements into questions.`
		glossary = s.glossaryEn
	} else if targetLang == "fr" {
		glossaryRules = "RÈGLES OBLIGATOIRES DU GLOSSAIRE :\n- Utilisez les entrées du glossaire uniquement lorsque le terme ou l'expression russe apparaît dans le texte source.\n- Si une entrée contient \"/\", choisissez exactement une seule option.\n- Ne copiez jamais dans la traduction des variantes du type \"A / B\".\n- Conservez le sujet et l'objet originaux de chaque phrase.\n"
		glossaryHeader = `GLOSSAIRE OBLIGATOIRE : Lorsque ces expressions russes apparaissent dans le texte source, utilisez la traduction correspondante. Ne les réinterprétez pas. Ne transformez pas les affirmations en questions.`
		glossary = s.glossaryFr
	}

	systemPrompt, ok := s.prompts[targetLang]
	if !ok {
		return "", fmt.Errorf("prompt not found for language: %s", targetLang)
	}

	injectedContext := glossaryRules + glossaryHeader
	for ru, trans := range glossary {
		if whitelist[ru] {
			injectedContext += fmt.Sprintf("- '%s' -> '%s'\n", ru, trans)
		}
	}

	userText := fmt.Sprintf("\n### ORIGINAL TEXT START ###\n%s\n### ORIGINAL TEXT END ###\n", text)

	systemPrompt += injectedContext
	payload := LLMPayload{
		Model: s.model,
		Messages: []LLMMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userText},
		},
		Temperature: 0,
		TopP:        1,
		MaxTokens:   2000,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	log.Info().
		Str("url", s.url).
		Str("model", s.model).
		Str("target_lang", targetLang).
		Msg("Sending request to LiteLLM")

	req, err := http.NewRequestWithContext(ctx, "POST", s.url+"/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-litellm-api-key", s.key)

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(body io.ReadCloser) {
		if closeErr := body.Close(); closeErr != nil {
			log.Error().Err(closeErr).Msg("failed to close response body")
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("litellm error: status %d", resp.StatusCode)
	}

	var result LLMResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Choices[0].Message.Content, nil
}
