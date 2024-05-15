package GPT

import (
	"encoding/json"
	"fmt"
	"os"
)

func marshallJsonRequestGPT(userText string) ([]byte, error) {
	const op = "./API/GPT/marshallJsonRequestGPT"
	var mes [2]message
	mes[0] = message{"system", "Сегодня ты программист и должен решить проблему из программирования."}
	mes[1] = message{"user", userText}
	req := requestGPT{
		ModelUri:          modelUri(),
		CompletionOptions: completionOptions{Stream: false, Temperature: 0.1, MaxTokens: "2000"},
		Messages:          mes,
	}

	reqJSON, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("Json encoding failed %s: %w", op, err)
	}

	return reqJSON, nil
}

func modelUri() string {
	return "gpt://" + os.Getenv("X_FOLDER_ID") + "/yandexgpt/latest"
}
