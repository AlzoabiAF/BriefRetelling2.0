package GPT

import (
	"encoding/json"
	"fmt"
)

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

func jsonUnmarshallResponse(resp []byte) (string, error) {
	const op = "./Api/GPT/jsonUnMarshallRequest"
	r := &responseGPT{}
	if err := json.Unmarshal(resp, r); err != nil {
		return "Ошибка сервера", fmt.Errorf("dont unmarshall json file: %s: %w", op, err)
	}
	return r.Result.Alternatives[0].Messages.Text, nil
}
