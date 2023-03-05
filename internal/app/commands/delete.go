package commands

import (
	"encoding/json"
	"fmt"

	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

// DeleteItemAsk returns item and ask delete confirmation
func DeleteItemAsk(itemid string, chatid int64, bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	// fing item again
	w := product.FindID(strconv.FormatInt(chatid, 10) + itemid)

	// creation message text
	msgt := fmt.Sprintf("Are you SURE you want to DELETE this Item?\n%s %s %s\n:(",
		w.Manufacture,
		w.Model,
		w.ItemID)

	// create message
	msg := tgbotapi.NewEditMessageCaption(chatid, update.CallbackQuery.Message.MessageID, msgt)

	// add buttons
	del, err := json.Marshal(CallBackButton{Command: "delete", ItemIDbyUser: w.UniqID[len(strconv.FormatInt(chatid, 10)):]})
	if err != nil {
		del = nil
	}

	nope, err := json.Marshal(CallBackButton{Command: "canceldelete", ItemIDbyUser: w.UniqID[len(strconv.FormatInt(chatid, 10)):]})
	if err != nil {
		del = nil
	}

	// add buttons
	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yes, I'm sure. Delete", string(del)),
			tgbotapi.NewInlineKeyboardButtonData("No. Keep it", string(nope)),
		),
	)
	msg.ReplyMarkup = &kb
	// send
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

// DeleteItem returns item info and new quantity of items in catalog. Deletes item.
func DeleteItem(itemid string, chatid int64, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// update message
	mess := "deleting..."
	msgr := tgbotapi.NewEditMessageCaption(chatid, update.CallbackQuery.Message.MessageID, mess)
	if _, err := bot.Send(msgr); err != nil {
		log.Panic(err)
	}

	// get item and delete
	w, qty := product.DeleteID(strconv.FormatInt(chatid, 10) + itemid)

	mess = fmt.Sprintf("%s %s %s\n%s %s\nWAS DELETED\n%d items in catalog",
		w.Manufacture,
		w.Model,
		w.ItemID,
		w.Color,
		w.Comments,
		qty)

	// send
	msgr = tgbotapi.NewEditMessageCaption(chatid, update.CallbackQuery.Message.MessageID, mess)
	if _, err := bot.Send(msgr); err != nil {
		log.Panic(err)
	}
}

// DeleteItem returns item info and new quantity of items in catalog. Deletes item.
func CancelDeleteItem(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// update message
	mess := "nothing was deleted"
	msgr := tgbotapi.NewEditMessageCaption(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, mess)
	if _, err := bot.Send(msgr); err != nil {
		log.Panic(err)
	}
}
