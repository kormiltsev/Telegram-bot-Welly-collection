package product

import (
	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"sort"
	"strconv"
	"strings"
	"time"
)

var DeleteList map[string]Welly
var WellyList map[string]Welly
var UsersList map[string]User
var Ws = UW{
	Users:   make(map[string]User),
	Wellyes: make(map[string]Welly),
	Deleted: make(map[string]Welly),
}

// sample empty item
func NewWelly() Welly {
	return Welly{
		UniqID:      "000",
		UserID:      "000",
		ItemID:      "000",
		Manufacture: "000",
		Model:       "000",
		Color:       "000",
		TitleFoto:   "none",
		AllFoto:     make([]string, 0),
		Comments:    "000",
	}
}
func AddNewItem(i string, w Welly) int {
	w.UniqID = i + strconv.FormatInt(time.Now().Unix(), 10)
	Ws.Wellyes[w.UniqID] = w
	SaveCatalog()
	return len(Ws.Wellyes)
}

func FindItems(uid string, ask string) []Welly {
	ws := make([]Welly, 0)
	for _, w := range Ws.Wellyes {
		if w.UserID == uid {
			if ask == "" {
				ws = append(ws, w)
			} else {
				if strings.Contains(w.Model, ask) || strings.Contains(w.Manufacture, ask) || strings.Contains(w.ItemID, ask) {
					ws = append(ws, w)
				}
			}
		}
	}
	sort.Slice(ws, func(i, j int) (less bool) {
		return ws[i].Manufacture < ws[j].Manufacture
	})
	return ws
}

func FindID(id string) Welly {
	return Ws.Wellyes[id]
}

func DeleteID(id string) (Welly, int) {
	w := FindID(id)
	delete(Ws.Wellyes, id)
	Ws.Deleted[id] = w
	SaveCatalog()
	return w, len(Ws.Wellyes)
}
func FindItemsDeleted(uid string, ask string) []Welly {
	ws := make([]Welly, 0)
	for _, w := range Ws.Deleted {
		if w.UserID == uid {
			if ask == "" {
				ws = append(ws, w)
			} else {
				if strings.Contains(w.Model, ask) || strings.Contains(w.Manufacture, ask) || strings.Contains(w.ItemID, ask) {
					ws = append(ws, w)
				}
			}
		}
	}
	sort.Slice(ws, func(i, j int) (less bool) {
		return ws[i].Manufacture < ws[j].Manufacture
	})
	return ws
}

func GetCatalog() (*UW, string) {

	WellyList = make(map[string]Welly)
	DeleteList = make(map[string]Welly)
	UsersList = make(map[string]User)

	UsersList["0"] = NewUser()
	WellyList["0"] = NewWelly()
	DeleteList["0"] = NewWelly()

	Ws = UW{
		Users:   UsersList,
		Wellyes: WellyList,
		Deleted: DeleteList,
	}
	return UploadCatalog(&Ws)
}

// func init() {
// 	WellyList = make(map[string]Welly)
// }

// func CheckUserExist(inputMessage *tgbotapi.Message) {
// 	//tgbotapi.update.Message.Chat.UserName
// 	u := NewUser()
// 	u.ID = inputMessage.Chat.ID
// 	u.NameF = inputMessage.Chat.FirstName
// 	u.NameL = inputMessage.Chat.LastName
// 	u.Username = inputMessage.Chat.UserName
// 	u.Type = inputMessage.Chat.Type

// 	if status := CheckUserDB(u); status != "ok" {
// 		log.Println(status)
// 	}
// }

// var UsersList map[string]User

// func CheckUserDB(u *User) string {
// 	if _, ok := UsersList[strconv.FormatInt(u.ID, 10)]; ok {
// 		return "ok"
// 	} else {
// 		UsersList[strconv.FormatInt(u.ID, 10)] = *u
// 		return "ok"
// 	}
// 	return "DB error: cant add user in User list"
// }
// func init() {
// 	UsersList = make(map[string]User)
// }
