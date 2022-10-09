package commands

import (
	"fmt"
	"log"
	"sort"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

func ShowCatalog(u *UserSpec, bot *tgbotapi.BotAPI) {
	mess := ""
	for _, w := range product.FindItems(u.Id, "") {
		mess = mess + fmt.Sprintf("%s %s %s\n",
			w.Manufacture,
			w.Model,
			w.ItemID)
	}
	mess = mess + "/menu"
	msg := tgbotapi.NewMessage(u.Id64, mess)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func ShowCatalogbyID(u *UserSpec, bot *tgbotapi.BotAPI) {
	mess := ""
	rang := make([]string, 0)
	cat := product.FindItems(u.Id, "")
	for _, w := range cat {
		rang = append(rang, w.ItemID)
	}
	sort.Strings(rang)
	for _, st := range rang {
		for _, w := range cat {
			if st == w.ItemID {
				mess = mess + fmt.Sprintf("%s %s %s\n",
					w.ItemID,
					w.Manufacture,
					w.Model)
			}
		}
	}
	mess = mess + "/menu"
	msg := tgbotapi.NewMessage(u.Id64, mess)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func ShowCatalogDeleted(u *UserSpec, bot *tgbotapi.BotAPI) {
	mess := ""
	for _, w := range product.FindItemsDeleted(u.Id, "") {
		mess = mess + fmt.Sprintf("%s %s %s\n",
			w.Manufacture,
			w.Model,
			w.ItemID)
	}
	mess = mess + "/menu"
	msg := tgbotapi.NewMessage(u.Id64, mess)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
