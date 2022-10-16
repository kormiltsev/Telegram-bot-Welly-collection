package commands

import (
	"log"
	"os"
	"strconv"

	"github.com/kormiltsev/tbot-welly/internal/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// func Menu1(bot *tgbotapi.BotAPI, update tgbotapi.Update, Users *Ids64) {
func SaveAndQuit(bot *tgbotapi.BotAPI) {
	godotenv.Load()
	mineID := os.Getenv("MYID")
	if mineID == "" {
		return
	}
	chatid64, err := strconv.ParseInt(mineID, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	//t := "@Collectionist_bot was stoped\n"
	catalogjson := product.CatalogAdres()
	//t := "Catalog-" + strconv.FormatInt(time.Now().Unix(), 10) + ".json"
	// data, _ := ioutil.ReadFile(catalogjson)
	// b := tgbotapi.FileBytes{Name: t, Bytes: data}
	// msg := tgbotapi.UploadData(mineID, b)
	var listMediaInput []interface{}

	//data, _ := ioutil.ReadFile(catalogjson)
	listMediaInput = append(listMediaInput, tgbotapi.NewInputMediaDocument(tgbotapi.FilePath(catalogjson)))

	msg := tgbotapi.NewMediaGroup(chatid64, listMediaInput)

	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
