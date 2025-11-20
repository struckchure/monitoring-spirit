package main

import (
	"github.com/spf13/cobra"
	ms "github.com/struckchure/monitoring-spirit"
)

var (
	model       string
	apiUrl      string
	apiKey      string
	apiProvider string
	promptType  string

	msService *ms.MsService
)

var rootCmd = &cobra.Command{
	Use:   "ms",
	Short: "Generate work reports based on your git commits with monitoring spirit.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		msService = ms.NewMsService(func(ac *ms.ApiConfig) *ms.ApiConfig {
			ac.Model = model
			ac.ApiUrl = apiUrl
			ac.ApiKey = apiKey
			ac.ApiProvider = ms.ApiProviderType(apiProvider)
			ac.PromptType = ms.PromptType(promptType)

			return ac
		})
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&model, "model", "m", "", "AI Model")
	rootCmd.PersistentFlags().StringVarP(&apiUrl, "api-url", "u", "", "Api Url")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", "", "Api Key")
	rootCmd.PersistentFlags().StringVarP(&apiProvider, "api-provider", "p", "ollama", "Api Provider (ollama, llmstudio)")
	rootCmd.PersistentFlags().StringVar(&promptType, "prompt-type", "default", "Prompt Type (default, technical, non-technical)")
}
