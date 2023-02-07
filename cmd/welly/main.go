package main

import (
	"flag"
	"log"
	"os"

	"github.com/kormiltsev/tbot-welly/internal/app/commands"
	"github.com/kormiltsev/tbot-welly/internal/product"

	env "github.com/caarlos0/env/v6"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var con envcon

type envcon struct {
	myToken string `env:"WELLY_TOKEN"`
	DBlink  string `env:"DATABASE_URL"`
}

// flags
func flags() {
	// in case of ENV
	err := env.Parse(con)
	if err != nil {
		log.Println(err)
	}
	// in case of .env
	godotenv.Load()
	pgurl := os.Getenv("DATABASE_URL")
	con.myToken = os.Getenv("TOKEN")

	// in case of flags
	{
		flag.StringVar(&pgurl, "pgurl", pgurl, "DATABASE_URL")
		tkn := flag.String("token", con.myToken, "TOKEN of Telegram bot")
		flag.Parse()
		con.myToken = *tkn
	}

	// push url to storage
	product.SetURL(pgurl)
}
func (c *SetByEnv) Environment() {
	err := env.Parse(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("got from env:", c)
}

func main() {
	//bot.Debug = true // debug status

	flags()
	bot, err := tgbotapi.NewBotAPI(con.myToken)
	if err != nil {
		panic(err)
	}
	defer commands.SaveAndQuit(bot)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// 24 hours and Telegram deletes all income messages
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	updates := bot.GetUpdatesChan(u)

	// upload catalog from file
	Ws, catok := product.GetCatalog()
	log.Println("Catalog status:", catok)

	// route updates
	for update := range updates {
		if update.CallbackQuery != nil {
			commands.HandleCallBack(bot, update, Ws)
		} else {
			commands.HandleUpdate(bot, update, Ws)
		}
	}
}
