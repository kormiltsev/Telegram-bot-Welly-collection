package commands

import (
	"log"
	"strconv"

	"github.com/kormiltsev/tbot-welly/internal/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mineID = "278284468"

// SaveAndQuit send whole catalog only to me in JSON file
// using to beckup
func SaveAndQuit(bot *tgbotapi.BotAPI) {
	chatid64, err := strconv.ParseInt(mineID, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	catalogjson := product.CatalogAdres()
	var listMediaInput []interface{}

	listMediaInput = append(listMediaInput, tgbotapi.NewInputMediaDocument(tgbotapi.FilePath(catalogjson)))

	msg := tgbotapi.NewMediaGroup(chatid64, listMediaInput)

	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
