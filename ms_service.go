package ms

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/samber/lo"
)

type MsService struct {
	apiConfig *ApiConfig
}

type FilterArgs struct {
	From  *string
	To    *string
	Since *time.Time
	Until *time.Time
	Email *string
}

func (m *MsService) Report(filter FilterArgs) ([]string, error) {
	filter.Since = lo.Ternary(filter.Since != nil, filter.Since, lo.ToPtr(time.Now().AddDate(0, 0, -7)))
	filter.Until = lo.Ternary(filter.Until != nil, filter.Until, lo.ToPtr(time.Now()))

	dir := os.Getenv("GIT_DIR")
	repo, err := git.PlainOpen(lo.Ternary(dir != "", dir, "."))
	if err != nil {
		return nil, err
	}

	logOptions := &git.LogOptions{
		From: lo.TernaryF(
			filter.From != nil,
			func() plumbing.Hash { return plumbing.NewHash(*filter.From) },
			func() plumbing.Hash { return plumbing.ZeroHash },
		),
		To: lo.TernaryF(
			filter.To != nil,
			func() plumbing.Hash { return plumbing.NewHash(*filter.To) },
			func() plumbing.Hash { return plumbing.ZeroHash },
		),
	}
	log, err := repo.Log(logOptions)
	if err != nil {
		return nil, err
	}

	commits := []string{}

	log.ForEach(func(c *object.Commit) error {
		if filter.Email != nil && c.Author.Email != *filter.Email {
			return nil
		}

		commits = append(commits, c.Message)
		return nil
	})

	aiServiceFn, ok := aiServiceMapping[m.apiConfig.ApiProvider]
	if !ok {
		return nil, fmt.Errorf("provider %s is not currently implemented", m.apiConfig.ApiProvider)
	}
	aiService := aiServiceFn(m.apiConfig)
	output, err := aiService.Generate(strings.Join(commits, "\n"))
	if err != nil {
		return nil, err
	}

	return []string{*output}, nil
}

func NewMsService(opts ...ApiConfigOpt) *MsService {
	apiConfig := &ApiConfig{}
	for _, opt := range opts {
		opt(apiConfig)
	}

	return &MsService{apiConfig: apiConfig}
}
