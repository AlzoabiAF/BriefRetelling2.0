package tgbot

import (
	"context"
	"fmt"
	"log"

	gpt "BriefRetelling2.0/API/GPT"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func sendMenu(chatId int64) error {
	progLanguage = ""
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

func sendMessageResultGpt(text string, user *tgbotapi.User, message *tgbotapi.Message) {
	res, err := gpt.GPT(progLanguage, text)

	if err != nil {
		log.Printf("ID: %d User: %s %s %s Error: %s\n", user.ID, user.UserName, user.LastName, user.FirstName, err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, res)
	msg.ParseMode = tgbotapi.ModeMarkdown

	_, err = bot.Send(msg)
	if err != nil {
		log.Printf("ID: %d User: %s %s %s Error: %s\n", user.ID, user.UserName, user.LastName, user.FirstName, err)
	}

}

func sendMessageError(id int64) error {
	res := "<b>Вам необходимо было отправить тестовое задание или фотографию.</b>"

	msg := tgbotapi.NewMessage(id, res)
	msg.ParseMode = tgbotapi.ModeHTML
	_, err := bot.Send(msg)

	return err

}
