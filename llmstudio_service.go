package ms

import (
	"github.com/samber/lo"
	"resty.dev/v3"
)

type LlmStudioService struct {
	apiConfig *ApiConfig
	client    *resty.Client
}

func (l *LlmStudioService) Generate(input string) (*string, error) {
	res, err := l.client.R().
		SetBody(LlmStudioGenerateRequest{
			Model: l.apiConfig.Model,
			Messages: []LlmStudioGenerateRequestMessage{
				{
					Role:    LlmStudioGenerateRequestRoleTypeSystem,
					Content: aiPromptMapping[PromptTypeDefault],
				},
				{
					Role:    LlmStudioGenerateRequestRoleTypeUser,
					Content: input,
				},
			},
			Stream: false,
		}).
		SetResult(&LlmStudioGenerateResponse{}).
		Post("v1/chat/completions")
	if err != nil {
		return nil, err
	}
	response := res.Result().(*LlmStudioGenerateResponse)

	return &response.Choices[0].Message.Content, nil
}

func NewLlmStudioService(opts ...ApiConfigOpt) AiService {
	apiConfig := &ApiConfig{}
	for _, opt := range opts {
		opt(apiConfig)
	}

	apiConfig.Model = lo.Ternary(lo.IsEmpty(apiConfig.Model), "qwen/qwen3-vl-4b", apiConfig.Model)
	apiConfig.ApiUrl = lo.Ternary(lo.IsEmpty(apiConfig.ApiUrl), "http://localhost:1234", apiConfig.ApiUrl)

	client := resty.New().SetBaseURL(apiConfig.ApiUrl)

	return &LlmStudioService{
		apiConfig: apiConfig,
		client:    client,
	}
}
