package translation

import (
	"Derzhavnaya/internal/config"
	"bytes"
	"encoding/json"
	"fmt"

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
	url     string
	key     string
	model   string
	prompts map[string]string
	client  *http.Client
}

func NewTranslationService(cfg *config.Config) Translator {
	timeout := time.Duration(cfg.Translation.TimeoutSec) * time.Second
	return &LiteLLMClient{
		url:     cfg.Translation.LiteLLMUrl,
		key:     cfg.Translation.LiteLLMKey,
		client:  &http.Client{Timeout: timeout},
		model:   cfg.Translation.ModelName,
		prompts: cfg.Translation.Prompts,
	}
}

func (s *LiteLLMClient) Translate(ctx context.Context, text string, targetLang string) (string, error) {
	prompt, ok := s.prompts[targetLang]
	if !ok {
		return "", fmt.Errorf("prompt not found for language: %s", targetLang)
	}

	payload := LLMPayload{
		Model: s.model,
		Messages: []LLMMessage{
			{Role: "system", Content: prompt},
			{Role: "user", Content: text},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	log.Info().Str("url", s.url).Str("key", s.key).Str("payload", string(jsonData)).Msg("Sending request to LiteLLM")
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
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("litellm error: status %d", resp.StatusCode)
	}

	var result LLMResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Choices[0].Message.Content, nil
}
