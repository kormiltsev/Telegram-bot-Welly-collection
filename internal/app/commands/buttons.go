package commands

import (
	"encoding/json"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

type CallBackButton struct {
	Command      string `json:"command,omitempty"`
	ItemIDbyUser string `json:"item_id_by_user,omitempty"`
	Ask          string `json:"ask,omitempty"`
}

// Main hendler:
func HandleCallBack(bot *tgbotapi.BotAPI, update tgbotapi.Update, Ws *product.UW) {
	defer func() { //Panic
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	u, ok := Npc.Amigos[strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)]
	if !ok {
		log.Println("error with user recognition")
		return
	}

	// switcher =====================
	parsedData := CallBackButton{}
	json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
	switch parsedData.Command {
	case "askwithphoto":
		FindItemWithPhoto(bot, update)
		return
	case "photos":
		SendGallery(&u, bot, update)
		return
	case "deleteask":
		DeleteItemAsk(&u, bot, update)
		return
	case "delete":
		DeleteItem(&u, bot, update)
		return
	default:
		log.Printf("wrong Task in Button")
	}
	//=================================

	return
}

func SendGallery(u *UserSpec, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// s := strings.Split(update.Message.Text, "_")
	// log.Println(s)
	data := CallBackButton{}
	json.Unmarshal([]byte(update.CallbackQuery.Data), &data)
	ask := data.ItemIDbyUser

	w := product.FindID(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10) + ask)
	phots := make([]string, 0)
	if len(w.TitleFoto) >= 60 {
		phots = append(phots, w.TitleFoto)
	}

	phots = append(phots, w.AllFoto...)

	var listMediaPhotoInput []interface{}

	for i := 0; i < len(phots); i++ {
		if i == 10 {
			msg := tgbotapi.NewMediaGroup(update.CallbackQuery.Message.Chat.ID, listMediaPhotoInput)
			_, _ = bot.Send(msg)
			// if err != nil {
			// some tgbotapi error occures every time. need to check work with media group
			// log.Println(err)
			// }
			listMediaPhotoInput = listMediaPhotoInput[:1]
		}
		listMediaPhotoInput = append(listMediaPhotoInput, tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(phots[i])))
	}
	msg := tgbotapi.NewMediaGroup(update.CallbackQuery.Message.Chat.ID, listMediaPhotoInput)
	_, _ = bot.Send(msg)
	// if err != nil {
	// some tgbotapi error occures every time. need to check work with media group
	// log.Println(err)
	// }
}

// "fmt"
// "github.com/NautiloosGo/tbot-welly/internal/product"
// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// "log"

// func MenuButtons(inputMessage *tgbotapi.Message) {
// ?/?? ? msg.ReplyMarkup
// 	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("/add"),
// 			tgbotapi.NewKeyboardButton("/start"),
// 			tgbotapi.NewKeyboardButton("/catalog"),
// 		),
// 		// tgbotapi.NewKeyboardButtonRow(
// 		// 	tgbotapi.NewKeyboardButton("/delprod"),
// 		// tgbotapi.NewKeyboardButton("/get"),
// 		// tgbotapi.NewKeyboardButton("/editprod"),
// 		// ),
// 	)
// }

// var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
//     tgbotapi.NewInlineKeyboardRow(
//         tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
//         tgbotapi.NewInlineKeyboardButtonData("2", "2"),
//         tgbotapi.NewInlineKeyboardButtonData("3", "3"),
//     ),
//     tgbotapi.NewInlineKeyboardRow(
//         tgbotapi.NewInlineKeyboardButtonData("4", "4"),
//         tgbotapi.NewInlineKeyboardButtonData("5", "5"),
//         tgbotapi.NewInlineKeyboardButtonData("6", "6"),
//     ),
// )
// msg.ReplyMarkup = numericKeyboard

// вставляет в строку ввода имя бота и текст:
// exampleQuery := "hello, world"
// markup := tg.NewInlineKeyboardMarkup(
// 	tg.NewInlineKeyboardRow(
// 		tg.InlineKeyboardButton{
// 			Text:                         "Try it",
// 			SwitchInlineQueryCurrentChat: &exampleQuery,
// 		},
// 	),
// )
// reply := tg.NewMessage(msg.Chat.ID, "hello, world!")
// reply.ReplyMarkup = &markup
// bot.Send(reply)

//открыть и закрыть меню
// switch update.Message.Text {
// case "open":
// 	msg.ReplyMarkup = numericKeyboard
// case "close":
// 	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
// }
// if _, err := bot.Send(msg); err != nil {
// 	log.Panic(err)
// }
