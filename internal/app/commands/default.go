package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Default is service message (in case of error)
func Default(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Unknown command: "+inputMessage.Text)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
