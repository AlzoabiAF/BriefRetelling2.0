package GPT

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Языки

func GPT(programmingLanguage, task string) (string, error) {
	const op = "./API/GPT/gpt"
	userTextRequest := fmt.Sprintf("Напиши код на %s: %s", programmingLanguage, task)
	reqJson, err := marshallJsonRequestGPT(userTextRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", os.Getenv("URL_GPT"), bytes.NewReader(reqJson))
	if err != nil {
		return "", fmt.Errorf("couldn't make a request: %s: %w", op, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Api-Key "+os.Getenv("IAM_TOKEN"))

	client := &http.Client{}
	resq, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("couldn't send a request: %s: %s", op, err)
	}
	defer resq.Body.Close()

	body, err := io.ReadAll(resq.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %s: %s", op, err)
	}

	return jsonUnmarshallResponse(body)
}
