package tgbot

import (
	"bufio"
	"context"
	"log"
	"os"

	gpt "BriefRetelling2.0/API/GPT"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	// Menu texts
	languageSelection = "<b>Выберите язык программмирования</b>"

	// Button texts
	backButton = "Back"

	// Store bot screaming status
	progLanguage = ""
	bot          *tgbotapi.BotAPI

	// Keyboard layout for the first menu. One button, one row
	menuSelectionLanguage = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(gpt.Golang, gpt.Golang),
			tgbotapi.NewInlineKeyboardButtonData(gpt.Cpp, gpt.Cpp),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(gpt.Javascript, gpt.Javascript),
			tgbotapi.NewInlineKeyboardButtonData(gpt.Python, gpt.Python),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(gpt.Java, gpt.Java),
			tgbotapi.NewInlineKeyboardButtonData(gpt.Csharp, gpt.Csharp),
		),
	)

	// Keyboard layout for the second menu. Two buttons, one per row
	menuSendingRequest = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(backButton, backButton),
		),
	)
)

func Tg() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_BOT"))
	if err != nil {
		// Abort if something is wrong
		log.Panic(err)
	}

	// Set this to true to log all interactions with telegram servers
	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// `updates` is a golang channel which receives telegram updates
	updates := bot.GetUpdatesChan(u)

	// Pass cancellable context to goroutine
	go receiveUpdates(ctx, updates)

	// Tell the user the bot is online
	log.Println("Start listening for updates. Press enter to stop")

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()
}
