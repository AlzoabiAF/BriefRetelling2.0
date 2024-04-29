package tgbot

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
	for {
		reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		log.Println(reader[:len(reader)-2])
		if reader[:len(reader)-2] == "stopKey=1234" {
			log.Println("Stopping the application")
			cancel()
			os.Exit(0)
		}
	}
}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handleMessage(update.Message)
		break

	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
		break
	}
}

func handleMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	// Print to console
	log.Printf("ID: %d User: %s %s %s Text: %s",
		user.ID,
		user.UserName, user.LastName, user.FirstName,
		text[:checkingNumberOfSymbol(len(text))],
	)

	var err error
	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, text, user.FirstName)
	} else if len(text) > 0 {
		if progLanguage == "" {
			err = sendMenu(user.ID)
		} else {
			res, err := gpt.GPT(progLanguage, text)
			progLanguage = ""
			if err != nil {
				log.Printf("ID: %d User: %s %s %s Error: %s\n", user.ID, user.UserName, user.LastName, user.FirstName, err)
			}
			msg := tgbotapi.NewMessage(message.Chat.ID, res)
			msg.ParseMode = tgbotapi.ModeMarkdown
			_, err = bot.Send(msg)
			if err != nil {
				log.Printf("ID: %d User: %s %s %s Error: %s\n", user.ID, user.UserName, user.LastName, user.FirstName, err)
			}
			err = sendMenu(message.Chat.ID)
		}

	} else {
		// This is equivalent to forwarding, without the sender's name
		copyMsg := tgbotapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
		_, err = bot.CopyMessage(copyMsg)
	}

	if err != nil {
		log.Printf("ID: %d User: %s %s %s Error: %s\n", user.ID, user.UserName, user.LastName, user.FirstName, err)
	}
}

// When we get a command, we react accordingly
func handleCommand(chatId int64, command string, firstName string) error {
	var err error

	switch command {
	case "/start":
		progLanguage = ""
		msg := tgbotapi.NewMessage(chatId, initialGreeting(firstName))
		msg.ParseMode = tgbotapi.ModeHTML
		_, err = bot.Send(msg)

		if err != nil {
			break
		}
	case "/menu":
		progLanguage = ""
		err = sendMenu(chatId)
		break
	}

	return err
}

func handleButton(query *tgbotapi.CallbackQuery) {
	var text string

	markup := tgbotapi.NewInlineKeyboardMarkup()
	message := query.Message
	user := query.From

	if query.Data == gpt.Golang {
		text = generateTextMenuSendingRequest(gpt.Golang)
		log.Printf("ID: %d User: %s %s %s Button: %s", user.ID, user.UserName, user.LastName, user.FirstName, gpt.Golang)
		progLanguage = gpt.Golang
		markup = menuSendingRequest
	} else if query.Data == gpt.Cpp {
		text = generateTextMenuSendingRequest(gpt.Cpp)
		log.Printf("ID: %d User: %s %s %s Button: %s", user.ID, user.UserName, user.LastName, user.FirstName, gpt.Cpp)
		progLanguage = gpt.Cpp
		markup = menuSendingRequest
	} else if query.Data == gpt.Javascript {
		text = generateTextMenuSendingRequest(gpt.Javascript)
		log.Printf("ID: %d User: %s %s %s Button: %s", user.ID, user.UserName, user.LastName, user.FirstName, gpt.Javascript)
		progLanguage = gpt.Javascript
		markup = menuSendingRequest
	} else if query.Data == gpt.Python {
		text = generateTextMenuSendingRequest(gpt.Python)
		log.Printf("ID: %d User: %s %s %s Button: %s", user.ID, user.UserName, user.LastName, user.FirstName, gpt.Python)
		progLanguage = gpt.Python
		markup = menuSendingRequest
	} else if query.Data == gpt.Java {
		text = generateTextMenuSendingRequest(gpt.Java)
		log.Printf("ID: %d User: %s %s %s Button: %s", user.ID, user.UserName, user.LastName, user.FirstName, gpt.Java)
		progLanguage = gpt.Java
		markup = menuSendingRequest
	} else if query.Data == gpt.Csharp {
		text = generateTextMenuSendingRequest(gpt.Csharp)
		log.Printf("ID: %d User: %s %s %s Button: %s", user.ID, user.UserName, user.LastName, user.FirstName, gpt.Csharp)
		progLanguage = gpt.Csharp
		markup = menuSendingRequest
	} else if query.Data == backButton {
		progLanguage = ""
		log.Printf("ID: %d User: %s %s %s Button: %s", user.ID, user.UserName, user.LastName, user.FirstName, backButton)
		text = languageSelection
		markup = menuSelectionLanguage
	}

	callbackCfg := tgbotapi.NewCallback(query.ID, "")
	bot.Send(callbackCfg)

	// Replace menu text and keyboard
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func sendMenu(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, languageSelection)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = menuSelectionLanguage
	_, err := bot.Send(msg)
	return err
}

func generateTextMenuSendingRequest(progLang string) string {
	progLang = "<b>Вы выбрали " + progLang + "</b>\n\nПрошу предоставить текст вашей задачи, требующей решения."
	return progLang
}

func checkingNumberOfSymbol(count int) int {
	if count < 100 {
		return count
	} else {
		return 100
	}
}

func initialGreeting(firstName string) string {
	msg := fmt.Sprintf("<b>Привет, %s. </b>Я предназначен для генерации кода под ваши запросы.\n\nСначала вы выбираете язык программирования, а потом скидываете текст вашей задачи.\n\n<b>Для генерации запроса напишите</b> /menu", firstName)
	return msg
}
