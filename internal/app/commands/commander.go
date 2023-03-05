package commands

import (
	"log"
	"strconv"

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

// NewUser add new user
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

	// recognise user, create if not exists
	u, ok := Npc.Amigos[strconv.FormatInt(update.Message.Chat.ID, 10)]
	if !ok {
		NewUser(update.Message.Chat.ID)
		u = Npc.Amigos[strconv.FormatInt(update.Message.Chat.ID, 10)]
	}

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
			Menu(&u, bot)
			return
		}
	} else {
		// if no command check is there dialog in progress?
		if d := u.Params["dialog"].Value; d != "none" {
			Dialogue(&u, bot, update.Message)
		} else {
			// search text by model, manufacture or welly id
			FindItems(&u, bot, update.Message) // works as list
			return
		}
	}
	if update.Message == nil {
		return
	}

	_, ok = Ws.Users[u.Id]
	if !ok {
		product.CheckUserExist(update.Message)
	}
}
