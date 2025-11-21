package ms

type ApiProviderType string

const (
	ApiProviderTypeOpenAI    ApiProviderType = "openai"
	ApiProviderTypeOllama    ApiProviderType = "ollama"
	ApiProviderTypeLlmStudio ApiProviderType = "llmstudio"
)

type ApiConfig struct {
	Model       string
	ApiUrl      string
	ApiKey      string
	ApiProvider ApiProviderType
	PromptType  PromptType
}

type ApiConfigOpt func(*ApiConfig) *ApiConfig
