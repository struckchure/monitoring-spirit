package ms

import (
	"github.com/samber/lo"
	"resty.dev/v3"
)

type OllamaService struct {
	apiConfig *ApiConfig
	client    *resty.Client
}

func (o *OllamaService) ai(request OllamaGenerateRequest) (*OllamaGenerateResponse, error) {
	res, err := o.client.R().
		SetBody(request).
		SetResult(&OllamaGenerateResponse{}).
		Post("api/generate")
	if err != nil {
		return nil, err
	}
	response := res.Result().(*OllamaGenerateResponse)

	return response, nil
}

func (o *OllamaService) Summarize(input string) (*string, error) {
	response, err := o.ai(OllamaGenerateRequest{
		Model:  o.apiConfig.Model,
		System: commitSummarizerPrompt,
		Prompt: input,
	})
	if err != nil {
		return nil, err
	}

	return &response.Response, nil
}

func (o *OllamaService) Generate(input string) (*string, error) {
	response, err := o.ai(OllamaGenerateRequest{
		Model:  o.apiConfig.Model,
		System: aiPromptMapping[o.apiConfig.PromptType],
		Prompt: input,
	})
	if err != nil {
		return nil, err
	}

	return &response.Response, nil
}

func NewOllamaService(apiConfig *ApiConfig) AiService {
	apiConfig.PromptType = lo.Ternary(lo.IsEmpty(apiConfig.PromptType), PromptTypeDefault, apiConfig.PromptType)
	apiConfig.Model = lo.Ternary(lo.IsEmpty(apiConfig.Model), "gemma3:1b", apiConfig.Model)
	apiConfig.ApiUrl = lo.Ternary(lo.IsEmpty(apiConfig.ApiUrl), "http://localhost:11434", apiConfig.ApiUrl)

	client := resty.New().SetBaseURL(apiConfig.ApiUrl)

	return &OllamaService{
		apiConfig: apiConfig,
		client:    client,
	}
}
