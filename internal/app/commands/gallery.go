package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

// SendGallery returns all photo by id
func SendGallery(ask string, chatid int64, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// get items
	w := product.FindID(strconv.FormatInt(chatid, 10) + ask)
	phots := make([]string, 0)

	// if photo exists add to message
	if len(w.TitleFoto) >= 60 {
		phots = append(phots, w.TitleFoto)
	}

	phots = append(phots, w.AllFoto...)

	if len(phots) <= 1 {

		mess := fmt.Sprintf("%s %s %s\n%s %s\n\nNo more photos",
			w.Manufacture,
			w.Model,
			w.ItemID,
			w.Color,
			w.Comments)
		msg := tgbotapi.NewEditMessageCaption(chatid, update.CallbackQuery.Message.MessageID, mess)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
	var listMediaPhotoInput []interface{}

	for i := 0; i < len(phots); i++ {
		if i == 10 {
			msg := tgbotapi.NewMediaGroup(chatid, listMediaPhotoInput)
			bot.Send(msg)
			listMediaPhotoInput = listMediaPhotoInput[:1]
		}
		listMediaPhotoInput = append(listMediaPhotoInput, tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(phots[i])))
	}
	msg := tgbotapi.NewMediaGroup(chatid, listMediaPhotoInput)
	bot.Send(msg)
}
