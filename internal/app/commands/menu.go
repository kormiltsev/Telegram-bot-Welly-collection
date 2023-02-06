package commands

import (
	"time"

	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// func Menu1(bot *tgbotapi.BotAPI, update tgbotapi.Update, Users *Ids64) {
func Menu1(u *UserSpec, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	t := "loading..."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, t)
	a, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
	<-time.After(time.Second * 1)
	t = "Catalog"
	msgr := tgbotapi.NewEditMessageText(update.Message.Chat.ID, a.MessageID, t)
	if _, err = bot.Send(msgr); err != nil {
		log.Panic(err)
	}

}

func Menu(u *UserSpec, bot *tgbotapi.BotAPI) {
	// mess := fmt.Sprintf("Hello! Welcome to catalog!")
	// send local photo
	// data, _ := ioutil.ReadFile("pic.png")
	// file := tgbotapi.FileBytes{Name: "AQAD5LsxG5gzCEt-", Bytes: data}
	// send photo was upload before
	file := tgbotapi.FileID("AgACAgIAAxkDAAIFKGLhUvVhKqcrdTToDwPgQ3xvIUWxAAJ8vTEbDQ3wSt4fooC36sHLAQADAgADdwADKQQ")
	msg := tgbotapi.NewPhoto(u.Id64, file)
	msg.Caption = "NGGYU"
	//msg := tgbotapi.NewMessage(u.Id64, "To find item send name or item ID")
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
