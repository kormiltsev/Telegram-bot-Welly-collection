package product

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User struct {
	ID       int64  `json:"id"`
	NameF    string `json:"first_name"`
	NameL    string `json:"last_name"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Dialog   string `json:"dialog"`
	NewItem  Welly  `json:"new_item"`
}

func NewUser() User {
	return User{
		ID:       0,
		NameF:    "TEST",
		NameL:    "TESTOV",
		Username: "TESTOROSSA",
		Type:     "private",
		Dialog:   "",
		NewItem: Welly{
			UniqID:      "",
			UserID:      "",
			ItemID:      "",
			Manufacture: "",
			Model:       "",
			Color:       "",
			TitleFoto:   "",
			AllFoto:     make([]string, 0),
			Comments:    "",
		},
	}
}

func CheckUserExist(inputMessage *tgbotapi.Message) {
	//tgbotapi.update.Message.Chat.UserName
	u := NewUser()
	u.ID = inputMessage.Chat.ID
	u.NameF = inputMessage.Chat.FirstName
	u.NameL = inputMessage.Chat.LastName
	u.Username = inputMessage.Chat.UserName
	u.Type = inputMessage.Chat.Type
	u.Dialog = ""

	if status := CheckUserDB(u); status != "ok" {
		log.Println(status)
	}
}

func CheckUserDB(u User) string {
	if _, ok := Ws.Users[strconv.FormatInt(u.ID, 10)]; ok {
		return "already axist"
	} else {
		Ws.Users[strconv.FormatInt(u.ID, 10)] = u
		SaveCatalog()
		return "Added user"
	}
	return "DB error: cant add user in User list"
}
