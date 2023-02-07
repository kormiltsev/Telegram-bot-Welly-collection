package commands

import (
	"encoding/json"
	"fmt"
	"time"

	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

// DeleteItemAsk returns item and ask delete confirmation
func DeleteItemAsk(itemid string, chatid int64, bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	w := product.FindID(strconv.FormatInt(chatid, 10) + itemid)

	// creation message text
	msgt := fmt.Sprintf("Are you SURE you want to DELETE this Item?\n%s %s %s\n:(",
		w.Manufacture,
		w.Model,
		w.ItemID)

	// create message
	msg := tgbotapi.NewMessage(chatid, msgt)

	// add buttons
	del, err := json.Marshal(CallBackButton{Command: "delete", ItemIDbyUser: w.UniqID[len(strconv.FormatInt(chatid, 10)):]})
	if err != nil {
		del = nil
	}

	// add buttons
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes, I'm sure. Delete", string(del)),
		),
	)

	// send
	a, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}

	// wait 5 secs and replace "delete" button
	<-time.After(time.Second * 5)
	t := "Time out"
	msgr := tgbotapi.NewEditMessageText(chatid, a.MessageID, t)
	if _, err = bot.Send(msgr); err != nil {
		log.Panic(err)
	}
}

// DeleteItem returns item info and new quantity of items in catalog. Deletes item.
func DeleteItem(itemid string, chatid int64, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// get item and delete
	w, qty := product.DeleteID(strconv.FormatInt(chatid, 10) + itemid)

	// set default image for message
	file := tgbotapi.FileID("AgACAgIAAxkBAAIGomLrzy2KPw72xfdESzZ38rTCaBi7AALmwDEbmLVhSzsBlE2C-TvzAQADAgADcwADKQQ")

	// set foto from item if exists
	if len(w.TitleFoto) >= 60 {
		file = tgbotapi.FileID(w.TitleFoto)
	}

	// create message
	msg := tgbotapi.NewPhoto(chatid, file)
	msg.Caption = fmt.Sprintf("%s %s %s\n%s %s\nWAS DELETED\n%d items in catalog",
		w.Manufacture,
		w.Model,
		w.ItemID,
		w.Color,
		w.Comments,
		qty)

	// send
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
