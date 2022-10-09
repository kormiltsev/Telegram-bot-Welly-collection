package commands

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

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
// if update.Message.Photo != nil {
// 	Photo := update.Message.Photo[0]
// 	PhotoID := Photo.FileID
// 	log.Printf(PhotoID)

func SendFoto(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, Ws *product.UW) {
	data, _ := ioutil.ReadFile("pic.png")
	b := tgbotapi.FileBytes{Name: "image.png", Bytes: data}
	msg := tgbotapi.NewPhoto(inputMessage.Chat.ID, b)
	//file := tgbotapi.FileID("AgACAgIAAxkDAAIFKGLhUvVhKqcrdTToDwPgQ3xvIUWxAAJ8vTEbDQ3wSt4fooC36sHLAQADAgADdwADKQQ")
	msg.Caption = "Test"

	_, err := bot.Send(msg)

	if err != nil {
		log.Println(err)
	}

	// send gallepy up to 10 photos:
	file := tgbotapi.FileID("AgACAgIAAxkBAAIGAmLpQ12qYRKTL-4aT8f022NzKz_OAALevzEb-6pJS9Hg5i87VdjUAQADAgADeAADKQQ")
	file1 := tgbotapi.FileID("AgACAgIAAxkBAAIFI2LhTt-AAV6I6bnyPG5TcyAszPAAA-S7MRuYMwhL_nHCrUFz3K0BAAMCAANzAAMpBA")
	file2 := tgbotapi.FileID("AgACAgIAAxkBAAIGAmLpQ12qYRKTL-4aT8f022NzKz_OAALevzEb-6pJS9Hg5i87VdjUAQADAgADeAADKQQ")
	file3 := tgbotapi.FileID("AgACAgIAAxkBAAIFI2LhTt-AAV6I6bnyPG5TcyAszPAAA-S7MRuYMwhL_nHCrUFz3K0BAAMCAANzAAMpBA")

	var listMediaPhotoInput []interface{}
	listMediaPhotoInput = append(listMediaPhotoInput, tgbotapi.NewInputMediaPhoto(file))
	listMediaPhotoInput = append(listMediaPhotoInput, tgbotapi.NewInputMediaPhoto(file1))
	listMediaPhotoInput = append(listMediaPhotoInput, tgbotapi.NewInputMediaPhoto(file2))
	listMediaPhotoInput = append(listMediaPhotoInput, tgbotapi.NewInputMediaPhoto(file3))

	msg2 := tgbotapi.NewMediaGroup(inputMessage.Chat.ID, listMediaPhotoInput)
	_, err = bot.Send(msg2)
	if err != nil {
		log.Println(err)
	}
}
