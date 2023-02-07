package commands

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

func Dialogue(u *UserSpec, bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	if inputMessage.Photo != nil { // if is photo
		if u.Params["dialog"].Value == "title_foto" {
			photoID := "none"
			if inputMessage.Text != "" {
			} else {
				photo := inputMessage.Photo[0]
				photoID = photo.FileID
			}
			EditParam("title_foto", photoID, u)
			EditParam("dialog", "color", u)
			mess := fmt.Sprintf("Title photo is set. Next send me a color. Also you can add more photos (up to 10 in 1 message) any time untill item was saved")
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			return
		} else {
			phot := inputMessage.Photo[0]
			EditParam(fmt.Sprintf("photo_%s", phot.FileID), phot.FileID, u)
			return
		}
	}
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
		if inputMessage.Text != "" {
		} else {
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
	// case "all_foto":
	// 	if inputMessage.Text != "" {
	// 	} else {
	// 		phot := inputMessage.Photo[0]
	// 		EditParam(fmt.Sprintf("photo_%s", phot.FileID), phot.FileID, u)
	// 	}
	// 	EditParam("dialog", "color", u)
	// 	mess := fmt.Sprintf("next: color")
	// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
	// 	if _, err := bot.Send(msg); err != nil {
	// 		log.Panic(err)
	// 	}
	case "color":
		EditParam("color", inputMessage.Text, u)
		EditParam("dialog", "comments", u)
		mess := fmt.Sprintf("add some comments")
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	case "comments":
		file := tgbotapi.FileID("AgACAgIAAxkBAAIGAmLpQ12qYRKTL-4aT8f022NzKz_OAALevzEb-6pJS9Hg5i87VdjUAQADAgADeAADKQQ") // "Never Gonna Give You Up" picture
		i := 0
		tx := "none"
		for key, v := range u.Params {
			if strings.HasPrefix(key, "photo_") {
				i++
				file = tgbotapi.FileID(v.Value)
				tx = v.Value
			}
		}
		EditParam("comments", inputMessage.Text, u)
		EditParam("dialog", "checkandsaveitem", u)
		if len(u.Params["title_foto"].Value) >= 60 {
			file = tgbotapi.FileID(u.Params["title_foto"].Value)
		} else {
			u.Params["title_foto"] = Param{
				Title: "title_foto",
				Value: tx,
			}
		}
		msg := tgbotapi.NewPhoto(inputMessage.Chat.ID, file)
		msg.Caption = fmt.Sprintf("%s =Manufactrure\n%s =Model\n%s =WellyID\nColor: %s\nComments: %s\n+ %d photos",
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
