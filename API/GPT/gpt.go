package GPT

import "fmt"

type message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

// Языки
const Golang = "golang"
const Python = "python"
const Cpp = "c++"
const Java = "java"
const Javascript = "javascript"

func GPT(programmingLanguage, task string) (string, error) {
	return "", fmt.Errorf("")
}
