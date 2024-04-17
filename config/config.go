package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	yandexGPT
	http_server
	bot
}

type yandexGPT struct {
	iamToken  string `yaml:"iam_token"`
	gptFolder string `yaml:"GPTfolder"`
}

type http_server struct {
	address     string        `yaml:"address"`
	timeout     time.Duration `yaml:"timeout"`
	idleTimeout time.Duration `yaml:"idle_timeout"`
}

type bot struct {
	token string `yaml:"token"`
}

func MustLoadConfig() *Config {
	const op = "./config"
	configPath := os.Getenv("CONFIG_PATH")
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("env file does not exist: %s, %w", op, err)
	}

	var cfg Config

	return &cfg
}
