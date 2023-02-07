package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Menu is for unsigned funckions
func Menu(u *UserSpec, bot *tgbotapi.BotAPI) {
	file := tgbotapi.FileID("AgACAgIAAxkDAAIFKGLhUvVhKqcrdTToDwPgQ3xvIUWxAAJ8vTEbDQ3wSt4fooC36sHLAQADAgADdwADKQQ")
	msg := tgbotapi.NewPhoto(u.Id64, file)
	msg.Caption = "NGGYU\n/add - add new item\n/start - info\n/catalog - list of all items"
	//msg := tgbotapi.NewMessage(u.Id64, "To find item send name or item ID")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/add"),
			tgbotapi.NewKeyboardButton("/start"),
			tgbotapi.NewKeyboardButton("/catalog"),
		),
	)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
