package product

import (
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

// NewWelly returns sample empty item
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

// AddNewItem add item to memory and push to file
func AddNewItem(i string, w Welly) int {
	w.UniqID = i + strconv.FormatInt(time.Now().Unix(), 10)
	Ws.Wellyes[w.UniqID] = w
	SaveCatalog()
	return len(Ws.Wellyes)
}

// FindItems returns array of items contains string (by Model, Manufacture or ID)
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

// FindID returns item by ID and returns item
func FindID(id string) Welly {
	return Ws.Wellyes[id]
}

// DeleteID move item from catalog to Delet catalog and returns item
func DeleteID(id string) (Welly, int) {
	w := FindID(id)
	delete(Ws.Wellyes, id)
	Ws.Deleted[id] = w
	SaveCatalog()
	return w, len(Ws.Wellyes)
}

// FindItemsDeleted returns list of item deleted
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

// GetCatalog create catalog
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
