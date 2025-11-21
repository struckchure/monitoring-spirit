package ms

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/glamour"
)

func RenderMarkdown(input string) string {
	out, err := glamour.Render(input, "dark")
	if err != nil {
		log.Fatal(err)
	}

	return out
}

type ChunkConfig struct {
	CommitsPerChunk int
	OverlapSize     int
	FinalSummary    bool
}

func ChunkAndProcess(commits []string, cfg ChunkConfig, aiService AiService) []string {
	var summaries []string

	for i := 0; i < len(commits); i += cfg.CommitsPerChunk {
		end := min(i+cfg.CommitsPerChunk, len(commits))

		// Add overlap from previous chunk
		start := i
		if i > 0 && cfg.OverlapSize > 0 {
			start = i - cfg.OverlapSize
		}

		chunk := commits[start:end]
		summary, err := aiService.Summarize(strings.Join(chunk, "\n"))
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(RenderMarkdown(*summary))
		summaries = append(summaries, *summary)
	}

	return summaries
}
