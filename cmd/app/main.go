package main

import (
	"log"

	"BriefRetelling2.0/tgbot"

	// "BriefRetelling2.0/config"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("not found .env file")
	}

	tgbot.Tg()
}
