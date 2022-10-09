package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kormiltsev/tbot-welly/internal/app/commands"
	"github.com/kormiltsev/tbot-welly/internal/product"

	//".internal/services/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	myToken := os.Getenv("TOKEN")
	if myToken == "" {
		log.Printf("input telegram bot unique token:") //in case no .env file, ask for TOKEN in terminal
		fmt.Fscan(os.Stdin, &myToken)
	} // as alternative you can start app in terminal with: TOKEN="dfgdfg" go run main.go
	bot, err := tgbotapi.NewBotAPI(myToken)
	if err != nil {
		log.Panic(err)
	}
	defer commands.SaveAndQuit(bot)

	//bot.Debug = true // debug status

	log.Printf("Authorized on account %s", bot.Self.UserName)
	// 24 hours and Telegram deletes all income messages
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	updates := bot.GetUpdatesChan(u)

	Ws, catok := product.GetCatalog() // Catalog from file
	log.Printf("Catalog: %s", catok)

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

	for update := range updates {
		//log.Println(update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text)
		if update.CallbackQuery == nil { //if not button with CallBack, then income print in terminal and starts hendler
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		commands.HandleUpdate(bot, update, Ws)
	}
}

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
// 			"id":278284468,
// 			"is_bot":false,
// 			"first_name":"Artem",
// 			"last_name":"Kormiltsev",
// 			"username":"Nautiloos_T",
// 			"language_code":"ru"},
// 		"message":{
// 			"message_id":1507,
// 			"from":{
// 				"id":5438764829,
// 				"is_bot":true,
// 				"first_name":"Collectionist",
// 				"username":"collectionist_bot"},
// 			"chat":{
// 				"id":278284468,
// 				"first_name":"Artem",
// 				"last_name":"Kormiltsev",
// 				"username":"Nautiloos_T",
// 				"type":"private"},
// 			"date":1659452398,
// 			"text":"1/2\nAudi A6 1123\nG Cabrio",
// 			"reply_markup":{
// 				"inline_keyboard":[[{
// 						"text":"Next item",
// 						"callback_data":"{\"Task\":\"nextitem\",\"Word\":\"Aud\",\"Index\":1}"}]]}},
// 		"chat_instance":"-6024425003900926239",
// 		"data":"{\"Task\":\"nextitem\",\"Word\":\"Aud\",\"Index\":1}"}}]}
