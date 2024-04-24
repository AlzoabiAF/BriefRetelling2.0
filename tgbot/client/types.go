package client

import "net/http"

type Client struct {
	host     string
	basePath string
	client   http.Client
}

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}
