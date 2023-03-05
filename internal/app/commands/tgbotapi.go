package commands

// file := tgbotapi.FileID("AgACAgIAAxkDAAIFKGLhUvVhKqcrdTToDwPgQ3xvIUWxAAJ8vTEbDQ3wSt4fooC36sHLAQADAgADdwADKQQ")

// reminder:
// update.Message.Chat.UserName
// update.Message.Chat.ID
// update.Message.Text
// update.Message.Photo
//tgbotapi.update.Message.Chat.UserName
//inputMessage.Chat.UserName
//update.CallbackQuery.Data
//update.CallbackQuery.Message.Chat.ID
//update.CallbackQuery.Message.MessageID

// 	//tgbotapi.update.Message.Chat.UserName
// 	u := NewUser()
// 	u.ID = inputMessage.Chat.ID
// 	u.NameF = inputMessage.Chat.FirstName
// 	u.NameL = inputMessage.Chat.LastName
// 	u.Username = inputMessage.Chat.UserName
// 	u.Type = inputMessage.Chat.Type

//msg.ReplyMarkp = tgbotapi.NewRemoveKeybard(true)

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

// 	//check for photo ?
// // update.Message.Photo
// if update.Message.Photo != nil {
// 	Photo := update.Message.Photo[0]
// 	PhotoID := Photo.FileID
// 	log.Printf(PhotoID)

// 	//check for Document ?
// if update.Message.Document != nil {
// 	WhatTheDocument(bot, update)
// 	return
// }

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

// Type of responder:
// Endpoint: getUpdates, response: {"ok":true,"result":[{
// "update_id":
// "message":{
// 		"message_id":
// 		"from":{"
// 			"chat":{"
// 				"date":1658932959,
// 					"photo":[
// 						{"file_id":"AgACAgIAAxkBAAIFI2LhTt-AAV6I6bnyPG5TcyAszPAAA-S7MRuYMwhL_nHCrUFz3K0BAAMCAANzAAMpBA","file_unique_id":"AQAD5LsxG5gzCEt4","file_size":842,"width":90,"height":42},
// 						{"file_id":"AgACAgIAAxkBAAIFI2LhTt-AAV6I6bnyPG5TcyAszPAAA-S7MRuYMwhL_nHCrUFz3K0BAAMCAANtAAMpBA","file_unique_id":"AQAD5LsxG5gzCEty","file_size":10333,"width":320,"height":148},
// 						{"file_id":"AgACAgIAAxkBAAIFI2LhTt-AAV6I6bnyPG5TcyAszPAAA-S7MRuYMwhL_nHCrUFz3K0BAAMCAAN4AAMpBA","file_unique_id":"AQAD5LsxG5gzCEt9","file_size":42599,"width":800,"height":369},
// 						{"file_id":"AgACAgIAAxkBAAIFI2LhTt-AAV6I6bnyPG5TcyAszPAAA-S7MRuYMwhL_nHCrUFz3K0BAAMCAAN5AAMpBA","file_unique_id":"AQAD5LsxG5gzCEt-","file_size":74669,"width":1280,"height":591}]}}]}

// Endpoint: getUpdates, response: {"ok":true,"result":[{
// 	"update_id":935505784,
// 	"callback_query":{
// 		"id":"1195222691715899675",
// 		"from":{
// 			"id":0000000000,
// 			"is_bot":false,
// 			"first_name":"A",
// 			"last_name":"Kormiltsev",
// 			"username":"Nau",
// 			"language_code":"en"},
// 		"message":{
// 			"message_id":1507,
// 			"from":{
// 				"id":5438764829,
// 				"is_bot":true,
// 				"first_name":"Collectionist",
// 				"username":"collectionist_bot"},
// 			"chat":{
// 				"id":0000000000,
// 				"first_name":"A",
// 				"last_name":"Kormiltsev",
// 				"username":"Nau",
// 				"type":"private"},
// 			"date":1659452398,
// 			"text":"1/2\nAudi A6 1123\nG Cabrio",
// 			"reply_markup":{
// 				"inline_keyboard":[[{
// 						"text":"Next item",
// 						"callback_data":"{\"Task\":\"nextitem\",\"Word\":\"Aud\",\"Index\":1}"}]]}},
// 		"chat_instance":"-6024425003900926239",
// 		"data":"{\"Task\":\"nextitem\",\"Word\":\"Aud\",\"Index\":1}"}}]}

// if update.CallbackQuery != nil {
// in case of JSON:
//=================================
// parsedData := NextItemButtinCarusel{}
// json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
// switch parsedData.Task {
// case "nextitem":
// 	FindNextItemCarusel(bot, &parsedData, update)
// 	log.Printf("Next page")
// default:
// 	log.Printf("wrong Task in Button")
// }
//=================================

// 	// just text
// 	switch update.CallbackQuery.Data {
// 	case "deletemsg":
// 		log.Printf("xxx")
// 	default:
// 		FindItemWithPhoto(bot, update)
// 		log.Printf("Show with photo")
// 	}
// 	// =========

// 	//check for photo ?
// // update.Message.Photo
// if update.Message.Photo != nil {
// 	Photo := update.Message.Photo[0]
// 	PhotoID := Photo.FileID
// 	log.Printf(PhotoID)

// other files:===================
//RequestFileData{

// file := tgbotapi.FilePath("tests/image.jpg")
// file := tgbotapi.FileID("AgACAgIAAxkDAALesF8dCjAAAa_…")
// file := tgbotapi.FileURL("https://i.imgur.com/unQLJIb.jpg")

// var reader io.Reader
// file := tgbotapi.FileReader{
//     Name: "image.jpg",
//     Reader: reader,
// }

// var data []byte
// file := tgbotapi.FileBytes{
//     Name: "image.jpg",
//     Bytes: data,
// }
//}
// ===============================
