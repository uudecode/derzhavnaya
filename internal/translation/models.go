package translation

type LLMMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMPayload struct {
	Model    string       `json:"model"`
	Messages []LLMMessage `json:"messages"`
}

type LLMResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
