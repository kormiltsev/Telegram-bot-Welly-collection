package commands

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

// AddItem starts dialog with item creation
func AddItem(u *UserSpec, bot *tgbotapi.BotAPI, Ws *product.UW) {

	log.Println("add begin: ", u.Params["dialog"])
	EditParam("dialog", "manufacture", u)
	mess := "Lets start from manufacture name. \nSend me like Ford or Audi"

	msg := tgbotapi.NewMessage(u.Id64, mess)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

}

// Dialogue is collect data of new item vie several messages
func Dialogue(u *UserSpec, bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	// photo accepts at all steps
	if inputMessage.Photo != nil {
		// if title photo expected
		if u.Params["dialog"].Value == "title_foto" {
			photoID := "none"
			if inputMessage.Text != "" {
			} else {
				photo := inputMessage.Photo[0]
				photoID = photo.FileID
			}
			EditParam("title_foto", photoID, u)
			EditParam("dialog", "color", u)
			mess := "Title photo is set. Next send me a color. Also you can add more photos (up to 10 in 1 message) any time untill item was saved"
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			return
			// if photo is not title
		} else {
			phot := inputMessage.Photo[0]
			EditParam(fmt.Sprintf("photo_%s", phot.FileID), phot.FileID, u)
			return
		}
	}
	// what parameter we are expecting
	switch u.Params["dialog"].Value {
	case "manufacture":
		EditParam("manufacture", inputMessage.Text, u)
		EditParam("dialog", "model", u)
		mess := fmt.Sprintf("next: model name.")
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "model":
		EditParam("model", inputMessage.Text, u)
		EditParam("dialog", "itemid", u)
		mess := fmt.Sprintf("next: Welly model ID number.")
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "itemid":
		EditParam("itemid", inputMessage.Text, u)
		EditParam("dialog", "title_foto", u)
		mess := fmt.Sprintf("next: 1 title foto (or [no] to skip)")
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "title_foto":
		photoID := "none"
		if inputMessage.Text == "" {
			photo := inputMessage.Photo[0]
			photoID = photo.FileID
		}
		EditParam("title_foto", photoID, u)
		EditParam("dialog", "color", u)
		mess := fmt.Sprintf("Next send me a color. Also you can add more photos (up to 10 in 1 message) any time untill item was saved")
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "color":
		EditParam("color", inputMessage.Text, u)
		EditParam("dialog", "comments", u)
		mess := fmt.Sprintf("add some comments")
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		// last parameter, show new item data
	case "comments":
		msg := tgbotapi.NewPhoto(inputMessage.Chat.ID, tgbotapi.FilePath("./nofoto.png"))
		i := 0
		tx := "none"
		for key, v := range u.Params {
			if strings.HasPrefix(key, "photo_") {
				i++
				msg = tgbotapi.NewPhoto(inputMessage.Chat.ID, tgbotapi.FileID(v.Value))
				tx = v.Value
			}
		}

		EditParam("comments", inputMessage.Text, u)
		EditParam("dialog", "checkandsaveitem", u)

		if len(u.Params["title_foto"].Value) >= 60 {
			msg = tgbotapi.NewPhoto(inputMessage.Chat.ID, tgbotapi.FileID(u.Params["title_foto"].Value))
		} else {
			u.Params["title_foto"] = Param{
				Title: "title_foto",
				Value: tx,
			}
		}
		msg.Caption = fmt.Sprintf("%s (Manufactrure)\n%s (Model)\n%s (WellyID)\nColor: %s\nComments: %s\n+ %d photos",
			u.Params["manufacture"].Value,
			u.Params["model"].Value,
			u.Params["itemid"].Value,
			u.Params["color"].Value,
			u.Params["comments"].Value,
			i,
		)
		//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Ok"),
				tgbotapi.NewKeyboardButton("Dont save"),
			),
		)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		// approve data to save
	case "checkandsaveitem":
		if inputMessage.Text == "ok" || inputMessage.Text == "Ok" || inputMessage.Text == "OK" {
			qty := PushNewWellyToCatalog(u)
			EditParam("dialog", "none", u)
			// u.NewItem = product.NewWelly()
			mess := fmt.Sprintf("Congrats! %d items in catalog\n/add", qty)
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		} else {
			EditParam("dialog", "none", u)
			mess := fmt.Sprintf("nothing was saved\n/add")
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		for key, _ := range u.Params {
			if strings.HasPrefix(key, "photo_") {
				delete(u.Params, key)
			}
		}
	default:
		mess := fmt.Sprintf("Ooopssy. Try again.")
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// PushNewWellyToCatalog prepare and push to storage
func PushNewWellyToCatalog(u *UserSpec) int {
	ph := make([]string, 0)
	for key, str := range u.Params {
		if strings.HasPrefix(key, "photo_") {
			ph = append(ph, str.Value)
		}
	}
	var welly = product.Welly{
		UserID:      u.Id,
		ItemID:      u.Params["itemid"].Value,
		Manufacture: u.Params["manufacture"].Value,
		Model:       u.Params["model"].Value,
		Color:       u.Params["color"].Value,
		TitleFoto:   u.Params["title_foto"].Value,
		AllFoto:     ph,
		Comments:    u.Params["comments"].Value,
	}
	err := product.UploadRowPostgres(&welly)
	if err != nil {
		log.Println("error write to postgres: ", err)
	}
	return product.AddNewItem(u.Id, welly)
}
