package ms

import (
	"strings"

	"github.com/samber/lo"
	"resty.dev/v3"
)

type LlmStudioService struct {
	apiConfig *ApiConfig
	client    *resty.Client
}

func (l *LlmStudioService) ai(request LlmStudioGenerateRequest) (*LlmStudioGenerateResponse, error) {
	res, err := l.client.R().
		SetBody(request).
		SetResult(&LlmStudioGenerateResponse{}).
		Post("v1/chat/completions")
	if err != nil {
		return nil, err
	}
	response := res.Result().(*LlmStudioGenerateResponse)

	return response, nil
}

func (l *LlmStudioService) Summarize(input string) (*string, error) {
	response, err := l.ai(LlmStudioGenerateRequest{
		Model: l.apiConfig.Model,
		Messages: []LlmStudioGenerateRequestMessage{
			{
				Role:    LlmStudioGenerateRequestRoleTypeSystem,
				Content: commitSummarizerPrompt,
			},
			{
				Role:    LlmStudioGenerateRequestRoleTypeUser,
				Content: input,
			},
		},
		MaxTokens: 1000,
		Stream:    false,
	})
	if err != nil {
		return nil, err
	}

	return &response.Choices[0].Message.Content, nil
}

func (l *LlmStudioService) Generate(input string) (*string, error) {
	chunkConfig := ChunkConfig{
		CommitsPerChunk: 50,
		OverlapSize:     5,
		FinalSummary:    true,
	}
	summaries := ChunkAndProcess(strings.Split(input, "\n"), chunkConfig, l)

	response, err := l.ai(LlmStudioGenerateRequest{
		Model: l.apiConfig.Model,
		Messages: []LlmStudioGenerateRequestMessage{
			{
				Role:    LlmStudioGenerateRequestRoleTypeSystem,
				Content: l.apiConfig.Prompt,
			},
			{
				Role:    LlmStudioGenerateRequestRoleTypeUser,
				Content: strings.Join(summaries, "\n"),
			},
		},
		MaxTokens: 1000,
		Stream:    false,
	})
	if err != nil {
		return nil, err
	}

	return &response.Choices[0].Message.Content, nil
}

func NewLlmStudioService(apiConfig *ApiConfig) AiService {
	apiConfig.Model = lo.Ternary(lo.IsEmpty(apiConfig.Model), "qwen/qwen3-vl-4b", apiConfig.Model)
	apiConfig.ApiUrl = lo.Ternary(lo.IsEmpty(apiConfig.ApiUrl), "http://localhost:1234", apiConfig.ApiUrl)

	client := resty.New().SetBaseURL(apiConfig.ApiUrl)

	return &LlmStudioService{
		apiConfig: apiConfig,
		client:    client,
	}
}
