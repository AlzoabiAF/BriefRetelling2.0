package client

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

const (
	getUpdateMethod   = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New() *Client {
	return mustConfig()
}

func mustConfig() *Client {
	const op = "./tgbot/client/connect/mustConfig"

	configPath := os.Getenv("CONFIG_CLIENT")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("env file does not exist: %s", op)
	}

	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Dont read config file: %s: %s", op, err)
	}

	var cfg *Client

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("Dont read config file: %s: %s", op, err)
	}

	return cfg
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdateMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdateResponse

	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendMessage(charID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(charID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return fmt.Errorf("cant send message: %w", err)
	}

	return nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("cant do request: %w", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cant do request ")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
