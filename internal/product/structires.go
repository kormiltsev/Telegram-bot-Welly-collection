package product

//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Welly struct {
	UniqID      string   `json:"uniq_id"`
	UserID      string   `json:"user_id"`
	ItemID      string   `json:"item_id"`
	Manufacture string   `json:"manufacture"`
	Model       string   `json:"model"`
	Color       string   `json:"color"`
	TitleFoto   string   `json:"title_foto"`
	AllFoto     []string `json:"all_foto"`
	Comments    string   `json:"comments"`
}

type UW struct {
	Users   map[string]User
	Wellyes map[string]Welly
	Deleted map[string]Welly
}

type User struct {
	ID       int64  `json:"id"`
	NameF    string `json:"first_name"`
	NameL    string `json:"last_name"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Dialog   string `json:"dialog"`
	NewItem  Welly  `json:"new_item"`
}

type Confpg struct {
	Adress   string
	User     string
	Password string
	DB       string
}
