package ms

import (
	"github.com/samber/lo"
	"resty.dev/v3"
)

type OllamaService struct {
	apiConfig *ApiConfig
	client    *resty.Client
}

func (o *OllamaService) Generate(input string) (*string, error) {
	res, err := o.client.R().
		SetBody(OllamaGenerateRequest{
			Model:  o.apiConfig.Model,
			System: aiPromptMapping[o.apiConfig.PromptType],
			Prompt: input,
		}).
		SetResult(&OllamaGenerateResponse{}).
		Post("api/generate")
	if err != nil {
		return nil, err
	}
	response := res.Result().(*OllamaGenerateResponse)

	return &response.Response, nil
}

func NewOllamaService(opts ...ApiConfigOpt) AiService {
	apiConfig := &ApiConfig{}
	for _, opt := range opts {
		opt(apiConfig)
	}

	apiConfig.Model = lo.Ternary(lo.IsEmpty(apiConfig.Model), "gemma3:1b", apiConfig.Model)
	apiConfig.ApiUrl = lo.Ternary(lo.IsEmpty(apiConfig.ApiUrl), "http://localhost:11434", apiConfig.ApiUrl)

	client := resty.New().SetBaseURL(apiConfig.ApiUrl)

	return &OllamaService{
		apiConfig: apiConfig,
		client:    client,
	}
}
