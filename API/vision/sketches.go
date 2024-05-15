package vision

// import (
// 	"bytes"
// 	"encoding/base64"
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"
// )

// const path = "./API/vision/"
// const url = ""

// func Vision(r io.Reader) string {
// 	const op = path + "Vision"

// 	content, err := io.ReadAll(r)
// 	if err != nil {
// 		handlerErrors(op, err)
// 		return ""
// 	}

// 	body := createBody(base64.StdEncoding.EncodeToString(content))

// 	req, err := http.NewRequest(http.MethodPost, url, body)
// 	if err != nil {
// 		handlerErrors(op, err)
// 		return ""
// 	}

// 	client := http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println(creatingError(op, err))
// 		return ""
// 	}
// 	defer resp.Body.Close()

// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println(creatingError(op, err))
// 		return ""
// 	}

// 	var data response
// 	err = json.Unmarshal(bodyBytes, &resp)
// 	if err != nil {
// 		handlerErrors(op, err)
// 		return ""
// 	}

// 	// finish processing the response
// 	return ""
// }

// func createBody(content string) (io.Reader, error) {
// 	const op = path + "createBody"

// 	bodyData := body{
// 		mineTime: "image",
// 		language: []string{"*"},
// 		model:    "page",
// 		content:  content,
// 	}

// 	b, err := json.Marshal(&bodyData)
// 	if err != nil {
// 		// handlerErrors(op, err)
// 		return nil, err
// 	}

// 	return bytes.NewReader(b), nil
// }

// func creatingFilalText(r *response) string{
// 	const op = path + "creatingFinalText"

// 	var finalText string

// 	for _, block := range r.TextAnnotation.Blocks{
// 		for _, line := range block.Lines{
// 			for _, alternative := range line.Alternatives{
// 				finalText += alternative.Text
// 			}
// 		}
// 	}

// 	return finalText
// }
