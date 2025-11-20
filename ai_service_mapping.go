package ms

var aiServiceMapping = map[ApiProviderType]func(*ApiConfig) AiService{
	ApiProviderTypeOllama:    func(ac *ApiConfig) AiService { return NewOllamaService(func(*ApiConfig) *ApiConfig { return ac }) },
	ApiProviderTypeLlmStudio: func(ac *ApiConfig) AiService { return NewLlmStudioService(func(*ApiConfig) *ApiConfig { return ac }) },
}
