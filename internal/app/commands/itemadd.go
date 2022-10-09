package commands

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
	//"strconv"
)

func AddItem(u *UserSpec, bot *tgbotapi.BotAPI, Ws *product.UW) {

	log.Println("add begin: ", u.Params["dialog"])
	EditParam("dialog", "manufacture", u)
	mess := fmt.Sprintf("Lets start from manufacture name. \nSend me like Ford or Audi")

	msg := tgbotapi.NewMessage(u.Id64, mess)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

}
