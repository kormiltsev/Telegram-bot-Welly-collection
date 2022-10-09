package commands

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(u *UserSpec, bot *tgbotapi.BotAPI) {
	mess := fmt.Sprintf("Hello! Welcome to catalog!")
	msg := tgbotapi.NewMessage(u.Id64, mess)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/add"),
			tgbotapi.NewKeyboardButton("/start"),
			tgbotapi.NewKeyboardButton("/catalog"),
		),
		// tgbotapi.NewKeyboardButtonRow(
		// 	tgbotapi.NewKeyboardButton("/search"),
		// tgbotapi.NewKeyboardButton("/get"),
		// tgbotapi.NewKeyboardButton("/editprod"),
		// ),
	)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
