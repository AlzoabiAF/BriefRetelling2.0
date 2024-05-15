package vision

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const path = "./API/vision/"

func Vision(r io.Reader) (string, error) {
	const op = path + "Vision"

	content, err := io.ReadAll(r)
	if err != nil {
		return "", wrapError(err, op, "failed to read content")
	}

	body, err := createBody(base64.StdEncoding.EncodeToString(content))
	if err != nil {
		return "", wrapError(err, op, "failed to create body")
	}

	req, err := http.NewRequest(http.MethodPost, os.Getenv("URL_VISION"), body)
	if err != nil {
		return "", wrapError(err, op, "failed to create request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Api-Key "+os.Getenv("IAM_TOKEN_VISION"))
	req.Header.Set("x-folder-id", os.Getenv("X_FOLDER_ID"))
	req.Header.Set("x-data-logging-enabled", "true")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", wrapError(err, op, "failed to do request")
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", wrapError(err, op, "failed to read response body")
	}

	var respData response
	err = json.Unmarshal(bodyBytes, &respData)
	if err != nil {
		return "", wrapError(err, op, "failed to unmarshal response")
	}

	return respData.Result.TextAnnotation.FullText, nil
}

func createBody(content string) (io.Reader, error) {
	const op = path + "createBody"

	bodyData := body{
		MineTime:      "image",
		LanguageCodes: []string{"*"},
		Model:         "page",
		Content:       content,
	}

	b, err := json.Marshal(&bodyData)
	if err != nil {
		return nil, wrapError(err, op, "failed to marshal body data")
	}

	return bytes.NewReader(b), nil
}
