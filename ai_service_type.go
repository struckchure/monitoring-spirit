package ms

type AiService interface {
	Summarize(string) (*string, error)
	Generate(string) (*string, error)
}
