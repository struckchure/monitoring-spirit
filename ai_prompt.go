package ms

type PromptType string

var (
	PromptTypeDefault      PromptType = "default"
	PromptTypeTechnical    PromptType = "technical"
	PromptTypeNonTechnical PromptType = "non-technical"
)

var aiPromptMapping = map[PromptType]string{
	PromptTypeDefault: defaultPrompt,
}
