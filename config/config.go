package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	YandexConf
	Http_server
	Bot
}

type YandexConf struct {
	IAMtoken  string
	GptFolder string
}

type Http_server struct {
	Address      string
	Timeout      time.Duration
	Idle_timeout time.Duration
}

type Bot struct {
	Token string
}

func MustLoadConfig() *Config {
	const op = "./config"
	configPath := os.Getenv("CONFIG_PATH")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("env file does not exist: %s", op)
	}

	var cfg Config

	viper.SetConfigFile("config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Dont read config file: %s: %s", op, err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Dont unmarshal config file: %s: %s", op, err)
	}
	return &cfg
}
