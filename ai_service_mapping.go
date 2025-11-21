package ms

var aiServiceMapping = map[ApiProviderType]func(*ApiConfig) AiService{
	ApiProviderTypeOllama:    func(ac *ApiConfig) AiService { return NewOllamaService(ac) },
	ApiProviderTypeLlmStudio: func(ac *ApiConfig) AiService { return NewLlmStudioService(ac) },
}
