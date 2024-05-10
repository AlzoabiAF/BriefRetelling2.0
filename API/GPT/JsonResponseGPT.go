package GPT

import (
	"encoding/json"
	"fmt"
)

func jsonUnmarshallResponse(resp []byte) (string, error) {
	const op = "./Api/GPT/jsonUnMarshallResponse"
	r := &responseGPT{}
	if err := json.Unmarshal(resp, r); err != nil {
		return "Ошибка сервера", fmt.Errorf("dont unmarshall json file: %s: %w", op, err)
	}
	return r.Result.Alternatives[0].Messages.Text, nil
}
