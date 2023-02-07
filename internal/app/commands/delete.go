package commands

import (
	"encoding/json"
	"fmt"
	"time"

	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
	//"strings"
)

func DeleteItemAsk(u *UserSpec, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// s := strings.Split(inputMessage.Text, "_")
	data := CallBackButton{}
	json.Unmarshal([]byte(update.CallbackQuery.Data), &data)
	itemid := data.ItemIDbyUser

	w := product.FindID(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10) + itemid)
	msgt := fmt.Sprintf("Are you SURE you want to DELETE this Item?\n%s %s %s\n:(",
		w.Manufacture,
		w.Model,
		w.ItemID)

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, msgt)
	//msg.ReplyMarkp = tgbotapi.NewRemoveKeybard(true)
	del, err := json.Marshal(CallBackButton{Command: "delete", ItemIDbyUser: w.UniqID[len(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)):]})
	if err != nil {
		u = nil
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes, I'm sure. Delete", string(del)),
		),
	)

	a, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
	<-time.After(time.Second * 5)
	t := "Time out"
	msgr := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, a.MessageID, t)
	if _, err = bot.Send(msgr); err != nil {
		log.Panic(err)
	}
}

func DeleteItem(u *UserSpec, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// s := strings.Split(inputMessage.Text, "_")
	data := CallBackButton{}
	json.Unmarshal([]byte(update.CallbackQuery.Data), &data)
	itemid := data.ItemIDbyUser

	w, qty := product.DeleteID(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10) + itemid)

	file := tgbotapi.FileID("AgACAgIAAxkBAAIGomLrzy2KPw72xfdESzZ38rTCaBi7AALmwDEbmLVhSzsBlE2C-TvzAQADAgADcwADKQQ")
	if len(w.TitleFoto) >= 60 {
		file = tgbotapi.FileID(w.TitleFoto)
	}
	msg := tgbotapi.NewPhoto(update.CallbackQuery.Message.Chat.ID, file)
	msg.Caption = fmt.Sprintf("%s %s %s\n%s %s\nWAS DELETED\n%d items in catalog",
		w.Manufacture,
		w.Model,
		w.ItemID,
		w.Color,
		w.Comments,
		qty)
	//msg.ReplyMarkp = tgbotapi.NewRemoveKeybard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
