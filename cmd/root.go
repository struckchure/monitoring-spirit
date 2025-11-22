package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	ms "github.com/struckchure/monitoring-spirit"
)

var (
	model              string
	apiUrl             string
	apiKey             string
	apiProvider        string
	promptType         string
	promptsDirOVerride string

	msService *ms.MsService
)

func getPrompt(promptType string) (*string, error) {
	var file []byte
	var err error

	if !lo.IsEmpty(promptsDirOVerride) {
		file, err = os.ReadFile(path.Join(promptsDirOVerride, fmt.Sprintf("%s.md", promptType)))
		if err != nil {
			return nil, err
		}
	} else {
		file, err = ms.PromptsFS.ReadFile(path.Join("prompts", fmt.Sprintf("%s.md", promptType)))
		if err != nil {
			return nil, err
		}
	}

	return lo.ToPtr(string(file)), nil
}

var rootCmd = &cobra.Command{
	Use:   "ms",
	Short: "Generate work reports based on your git commits with monitoring spirit.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		prompt, err := getPrompt(promptType)
		if err != nil {
			log.Panicln(err)
			return
		}

		apiConfig := &ms.ApiConfig{
			Model:       model,
			ApiUrl:      apiUrl,
			ApiKey:      apiKey,
			ApiProvider: ms.ApiProviderType(apiProvider),
			Prompt:      *prompt,
		}
		msService = ms.NewMsService(apiConfig)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&model, "model", "m", "", "AI Model")
	rootCmd.PersistentFlags().StringVarP(&apiUrl, "api-url", "u", "", "Api Url")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", "", "Api Key")
	rootCmd.PersistentFlags().StringVarP(&apiProvider, "api-provider", "p", "ollama", "Api Provider (ollama, llmstudio)")
	rootCmd.PersistentFlags().StringVar(&promptType, "prompt-type", "default", "Prompt Type (default, technical, non-technical)")
	rootCmd.PersistentFlags().StringVar(&promptsDirOVerride, "prompt-dir", "", "Prompt Directory")
}
