package tgbot

import (
	"BriefRetelling2.0/tgbot/client"
	"os"
)

func tg() {

	// TODO: tgclietn

	tgToken := os.Getenv("TG_ BOT")
	_ = tgToken

	a := client.New()

	_ = a

	// TODO: fetcher

	// TODO: processor

	// TODO: consumer

}
