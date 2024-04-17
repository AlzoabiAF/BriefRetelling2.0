package GPT

import (
	"encoding/json"
	"fmt"
)

type requestGPT struct {
	Authorization string `json:"Authorization"`
	x_folder_id   string `json:"x-folder-id"`
}

func JsonRequest(token_iam, token_folder string) ([]byte, error) {
	const op = "./Api/GPT/request"
	req := requestGPT{Authorization: "Authorization: Bearer " + token_iam, x_folder_id: "x-folder-id: " + token_folder}
	reqJSON, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("Json encoding failed %w: %s", err, op)
	}
	return reqJSON, nil
}

func JsonResponse([]byte) {}

func RequestGPT(json []byte) {

}

func ResponseGPT() {

}
