package commands

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
)

type Chars struct {
	Amigos map[string]UserSpec // key = 'string' of user ID strconv.FormatInt(update.Message.Chat.ID, 10)
}

type UserSpec struct {
	Id     string
	Id64   int64
	Params map[string]Param
}

type Param struct {
	Title string
	Value string
}

var Amig = make(map[string]UserSpec)
var Npc = Chars{
	Amigos: Amig,
}

func NewUser(i64 int64) {
	us := UserSpec{
		Id:     strconv.FormatInt(i64, 10),
		Id64:   i64,
		Params: make(map[string]Param),
	}
	EditParam("dialog", "none", &us)
	Npc.Amigos[us.Id] = us
}

func EditParam(t, v string, u *UserSpec) {
	u.Params[t] = Param{Title: t, Value: v}
}

// Main hendler:
func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, Ws *product.UW) {
	defer func() { //Panic
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	// if buttons
	if update.CallbackQuery != nil {
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

		// just text
		switch update.CallbackQuery.Data {
		case "xxx":
			log.Printf("xxx")
		default:
			FindItemWithPhoto(bot, update)
			log.Printf("Show with photo")
		}
		// =========

		return
	}

	// recognise user
	u, ok := Npc.Amigos[strconv.FormatInt(update.Message.Chat.ID, 10)]
	if !ok {
		NewUser(update.Message.Chat.ID)
		u = Npc.Amigos[strconv.FormatInt(update.Message.Chat.ID, 10)]
		//Npc.Amigos[uid].Params["dialog"] = ""
	}
	log.Println("user ad start", &u)

	// 	//check for photo ?
	// // update.Message.Photo
	// if update.Message.Photo != nil {
	// 	Photo := update.Message.Photo[0]
	// 	PhotoID := Photo.FileID
	// 	log.Printf(PhotoID)

	// if update.Message.Document != nil {
	// 	WhatTheDocument(bot, update)
	// 	return
	// }

	// check for command
	if update.Message.IsCommand() {
		switch update.Message.Text {
		case "/start":
			EditParam("dialog", "none", &u)
			Start(&u, bot)
			return
		case "/catalog":
			EditParam("dialog", "none", &u)
			ShowCatalog(&u, bot)
			return
		case "/catalogid":
			EditParam("dialog", "none", &u)
			ShowCatalogbyID(&u, bot)
			return
		case "/catalogdeleted":
			EditParam("dialog", "none", &u)
			ShowCatalogDeleted(&u, bot)
			return
		case "/catalogjsonsendme":
			SaveAndQuit(bot)
			return
		case "/menu":
			EditParam("dialog", "none", &u)
			Menu(&u, bot)
			return
		case "/add":
			EditParam("dialog", "none", &u)
			AddItem(&u, bot, Ws)
			return
		default:
			EditParam("dialog", "none", &u)
			if strings.HasPrefix(update.Message.Text, "/showallphotos_") {
				SendGallery(&u, bot, update.Message)
				return
			}
			if strings.HasPrefix(update.Message.Text, "/deleteitem_") {
				DeleteItem(&u, bot, update.Message)
				return
			}
			Menu(&u, bot)
			return
		}
	} else {
		if d := u.Params["dialog"].Value; d != "none" {
			Dialogue(&u, bot, update.Message)
		} else {
			// search text by model, manufacture or welly id
			FindItems(&u, bot, update.Message) // works as list
			//Default(bot, update.Message)
			return
		}
	}
	if update.Message == nil { // If we got a message / skip non-message
		return
	}

	_, ok = Ws.Users[u.Id]
	if !ok {
		product.CheckUserExist(update.Message)
	}

	// if u := Ws.Users[strconv.FormatInt(update.Message.Chat.ID, 10)]; u.Dialog != "" {
	// 	c.Dialog(update.Message, &u)
	// }
	log.Println(u)
}
