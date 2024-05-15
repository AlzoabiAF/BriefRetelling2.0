package tgbot

import (
	"log"
	"net/http"
	"strings"

	gpt "BriefRetelling2.0/API/GPT"
	"BriefRetelling2.0/API/vision"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
	image := message.Photo
	if user == nil {
		return
	}

	// Print to console

	var err error
	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, text, user.FirstName)
	} else if progLanguage == "" {
		err = sendMenu(user.ID)
	} else if len(image) > 0 {
		fileID := image[len(image)-1].FileID
		fileURL, err := bot.GetFileDirectURL(fileID)
		if err != nil {
			log.Printf("ID: %d User: %s %s %s Error: %s\n", user.ID, user.UserName, user.LastName, user.FirstName, err)
		}

		resp, err := http.Get(fileURL)
		if err != nil {
			log.Printf("ID: %d User: %s %s %s Error: %s\n", user.ID, user.UserName, user.LastName, user.FirstName, err)
		}
		defer resp.Body.Close()

		text, err := vision.Vision(resp.Body)
		if err != nil {
			log.Printf("ID: %d User: %s %s %s Error: %s\n", user.ID, user.UserName, user.LastName, user.FirstName, err)
		}

		sendMessageResultGpt(text, user, message)
		err = sendMenu(message.Chat.ID)
	} else if len(text) > 0 {
		sendMessageResultGpt(text, user, message)
		err = sendMenu(message.Chat.ID)
	} else {
		sendMessageError(message.Chat.ID)
		if err != nil {
			log.Printf("ID: %d User: %s %s %s Error: %s\n", user.ID, user.UserName, user.LastName, user.FirstName, err)
		}
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
