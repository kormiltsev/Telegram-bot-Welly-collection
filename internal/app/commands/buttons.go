package commands

// "fmt"
// "github.com/NautiloosGo/tbot-welly/internal/product"
// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// "log"

// func MenuButtons(inputMessage *tgbotapi.Message) {
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
