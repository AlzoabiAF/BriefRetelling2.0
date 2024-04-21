package main

import (
	"BriefRetelling2.0/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("not found .env file")
	}

	cfg := config.MustLoadConfig()

	_ = cfg

	// TODO: GPT

	// TODO: bot

	// TODO: Run server
}
