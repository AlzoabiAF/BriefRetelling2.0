package GPT

const Golang = "Golang"
const Cpp = "C++"
const Javascript = "JavaScript"
const Python = "Python"
const Java = "Java"
const Csharp = "C#"

type message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type requestGPT struct {
	ModelUri          string            `json:"modelUri"`
	CompletionOptions completionOptions `json:"completionOptions"`
	Messages          [2]message        `json:"messages"`
}

type completionOptions struct {
	Stream      bool    `json:"stream"`
	Temperature float32 `json:"temperature"`
	MaxTokens   string  `json:"maxTokens"`
}

type responseGPT struct {
	Result result `json:"result"`
}

type result struct {
	Alternatives []alternative `json:"alternatives"`
	Usage        usage         `json:"usage"`
	ModelVersion string        `json:"modelVersion"`
}

type usage struct {
	InputTextTokens  string `json:"inputTextTokens"`
	CompletionTokens string `json:"completionTokens"`
	TotalTokens      string `json:"totalTokens"`
}

type alternative struct {
	Messages message `json:"message"`
	Status   string  `json:"status"`
}
