package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Start returns default info
func Start(u *UserSpec, bot *tgbotapi.BotAPI) {
	mess := "Hello! Welcome to welly 1:60 catalog!\nTo find item send model name, manufacture or item ID\n/start - info\n/add - add new item\n/catalog - list of items"
	msg := tgbotapi.NewMessage(u.Id64, mess)

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
