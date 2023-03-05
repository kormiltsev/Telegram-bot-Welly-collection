package commands

import (
	"encoding/json"
	"fmt"

	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

type NextItemButtinCarusel struct {
	Task  string
	Word  string
	Index int
}

// FindItems search item with model or manufacture name or ID
func FindItems(u *UserSpec, bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	ask := inputMessage.Text
	mess := ""
	i := 0
	// list all items found
	for _, w := range product.FindItems(u.Id, ask) {
		mess = mess + fmt.Sprintf("%s %s %s %s %s\n",
			w.ItemID,
			w.Manufacture,
			w.Model,
			w.Color,
			w.Comments)
		i++
	}
	// if not found
	if i <= 0 {
		mess := "Looking for " + ask + "\nbut nothing found"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		return
	}
	if i == 1 {
		FindItemWithPhoto(ask, inputMessage.Chat.ID, bot)
		return
	}

	messg := fmt.Sprintf("Looking for %s,\n found: %d\n", ask, i) + mess
	msgt := tgbotapi.NewMessage(inputMessage.Chat.ID, messg)

	ph, err := json.Marshal(CallBackButton{Command: "askwithphoto", Ask: ask})
	if err != nil {
		ph = nil
	}

	msgt.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Show with Photo", string(ph)),
		),
	)
	//msgt.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msgt); err != nil {
		log.Panic(err)
	}
}

// answer is 1 photo in 1 message per item found
func FindItemWithPhoto(ask string, chatid int64, bot *tgbotapi.BotAPI) {
	i := 0
	for _, w := range product.FindItems(strconv.FormatInt(chatid, 10), ask) {
		msg := tgbotapi.NewPhoto(chatid, tgbotapi.FilePath("./nofoto.png"))
		if len(w.TitleFoto) >= 60 {
			msg = tgbotapi.NewPhoto(chatid, tgbotapi.FileID(w.TitleFoto))
		}
		msg.Caption = fmt.Sprintf("%s %s %s\n%s %s\n",
			w.Manufacture,
			w.Model,
			w.ItemID,
			w.Color,
			w.Comments)
		ph, err := json.Marshal(CallBackButton{Command: "photos", ItemIDbyUser: w.UniqID[len(strconv.FormatInt(chatid, 10)):]})
		if err != nil {
			ph = nil
		}
		del, err := json.Marshal(CallBackButton{Command: "deleteask", ItemIDbyUser: w.UniqID[len(strconv.FormatInt(chatid, 10)):]})
		if err != nil {
			del = nil
		}
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("All Photos", string(ph)),
				tgbotapi.NewInlineKeyboardButtonData("Delete", string(del)),
			),
		)
		// send
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		i++
	}
	if i <= 0 {
		mess := "Looking for " + ask + "\nnot found"
		msg := tgbotapi.NewMessage(chatid, mess)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
