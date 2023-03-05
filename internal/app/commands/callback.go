package commands

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

type CallBackButton struct {
	Command      string `json:"command,omitempty"`
	ItemIDbyUser string `json:"item_id_by_user,omitempty"`
	Ask          string `json:"ask,omitempty"`
}

// Main hendler for CallBackQuery:
func HandleCallBack(bot *tgbotapi.BotAPI, update tgbotapi.Update, Ws *product.UW) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()
	// read button data
	data := CallBackButton{}
	json.Unmarshal([]byte(update.CallbackQuery.Data), &data)
	ask := data.Ask
	itemid := data.ItemIDbyUser

	// switcher =====================
	parsedData := CallBackButton{}
	json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
	switch parsedData.Command {
	case "askwithphoto":
		FindItemWithPhoto(ask, update.CallbackQuery.Message.Chat.ID, bot)
		return
	case "photos":
		SendGallery(itemid, update.CallbackQuery.Message.Chat.ID, bot, update)
		return
	case "deleteask":
		DeleteItemAsk(itemid, update.CallbackQuery.Message.Chat.ID, bot, update)
		return
	case "delete":
		DeleteItem(itemid, update.CallbackQuery.Message.Chat.ID, bot, update)
		return
	case "canceldelete":
		CancelDeleteItem(bot, update)
		return
	default:
		log.Printf("wrong Task in Button")
	}
	//=================================

	return
}
