package translation

type LLMMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMPayload struct {
	Model       string       `json:"model"`
	Messages    []LLMMessage `json:"messages"`
	Temperature float64      `json:"temperature,omitempty"`
	TopP        float64      `json:"top_p,omitempty"`
	MaxTokens   int          `json:"max_tokens,omitempty"`
}

type LLMResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
