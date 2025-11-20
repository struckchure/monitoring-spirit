package ms

type LlmStudioGenerateRequestRoleType string

const (
	LlmStudioGenerateRequestRoleTypeSystem LlmStudioGenerateRequestRoleType = "system"
	LlmStudioGenerateRequestRoleTypeUser   LlmStudioGenerateRequestRoleType = "user"
)

type LlmStudioGenerateRequestMessage struct {
	Role    LlmStudioGenerateRequestRoleType `json:"role"`
	Content string                           `json:"content"`
}

type LlmStudioGenerateRequest struct {
	Model     string                            `json:"model"`
	Messages  []LlmStudioGenerateRequestMessage `json:"messages"`
	Stream    bool                              `json:"stream"`
	MaxTokens int                               `json:"max_tokens"`
}

type LlmStudioGenerateResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role      LlmStudioGenerateRequestRoleType `json:"role"`
			Content   string                           `json:"content"`
			ToolCalls []interface{}                    `json:"tool_calls"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Stats struct {
	} `json:"stats"`
	SystemFingerprint string `json:"system_fingerprint"`
}
