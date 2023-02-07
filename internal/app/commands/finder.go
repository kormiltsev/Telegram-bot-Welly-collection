package commands

import (
	"encoding/json"
	"fmt"

	"log"
	"sort"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-welly/internal/product"
	//"strings"
)

type NextItemButtinCarusel struct {
	Task string
	// Chatid    int64
	// Messageid int
	Word  string
	Index int
	//List      []product.Welly
}

func FindItems(u *UserSpec, bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	ask := inputMessage.Text
	//var listMediaPhotoInput []interface{}
	mess := ""
	i := 0
	//onefoto := ""
	for _, w := range product.FindItems(u.Id, ask) {
		mess = mess + fmt.Sprintf("%s %s %s %s %s\n",
			w.ItemID,
			w.Manufacture,
			w.Model,
			w.Color,
			w.Comments)
		// if len(w.TitleFoto) >= 60 {
		// 	//listMediaPhotoInput = append(listMediaPhotoInput, tgbotapi.NewInputMediaPhoto(tgbotapi.FileID("AgACAgIAAxkBAAIFI2LhTt-AAV6I6bnyPG5TcyAszPAAA-S7MRuYMwhL_nHCrUFz3K0BAAMCAANzAAMpBA")))
		// 	onefoto = w.TitleFoto
		// 	listMediaPhotoInput = append(listMediaPhotoInput, tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(w.TitleFoto)))
		// }
		i++
		// if i%10 == 0 && i != 0 {
		// 	msgp := tgbotapi.NewMediaGroup(inputMessage.Chat.ID, listMediaPhotoInput)
		// 	if _, err := bot.Send(msgp); err != nil {
		// 		log.Panic(err)
		// 	}

		// 	listMediaPhotoInput = listMediaPhotoInput[:0]
		// }
	}
	if i <= 0 {
		mess := "Looking for " + ask + "\nbut nothing found"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		return
	}
	// if i == 1 {
	// 	file := tgbotapi.FileID(onefoto)
	// 	msg := tgbotapi.NewPhoto(inputMessage.Chat.ID, file)
	// 	msg.Caption = mess
	// 	if _, err := bot.Send(msg); err != nil {
	// 		log.Panic(err)
	// 	}
	// }
	// msgp := tgbotapi.NewMediaGroup(inputMessage.Chat.ID, listMediaPhotoInput)
	// if _, err := bot.Send(msgp); err != nil {
	// 	log.Panic(err)
	// }

	messg := fmt.Sprintf("Looking for %s,\n found: %d\n", ask, i) + mess
	msgt := tgbotapi.NewMessage(inputMessage.Chat.ID, messg)

	ph, err := json.Marshal(CallBackButton{Command: "askwithphoto", Ask: ask})
	if err != nil {
		ph = nil
	}

	msgt.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Show with Photo", string(ph)),
		),
	)
	//msgt.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msgt); err != nil {
		log.Panic(err)
	}
}

// answer is text
func FindItem(u *UserSpec, bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	ask := inputMessage.Text
	mess := "Looking for " + ask + "\n"
	i := 0
	for _, w := range product.FindItems(u.Id, ask) {
		//file := tgbotai.FileID("AgACAgIAAxkAAIFKGLhUVhKqcrdTToDwPgQ3xvIUWxAAJ8vTEbDQ3wSt4fooC36sHLAQADAgADdwADKQQ")
		if len(w.TitleFoto) <= 6 || w.TitleFoto == "" {
			mess = fmt.Sprintf("%s %s %s\n%s %s",
				w.Manufacture,
				w.Model,
				w.ItemID,
				w.Color,
				w.Comments)
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
			//msg.ReplyMarkp = tgbotapi.NewRemoveKeybard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		} else {
			file := tgbotapi.FileID(w.TitleFoto)
			msg := tgbotapi.NewPhoto(inputMessage.Chat.ID, file)
			msg.Caption = fmt.Sprintf("%s %s %s\n%s %s",
				w.Manufacture,
				w.Model,
				w.ItemID,
				w.Color,
				w.Comments)
			//msg.ReplyMarkp = tgbotapi.NewRemoveKeybard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		i++
	}
	if i <= 0 {
		mess := "Looking for " + ask + "\nnot found"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// answer is 1 photo in 1 message per item found
func FindItemWithPhoto(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// ask := update.CallbackQuery.Data
	data := CallBackButton{}
	json.Unmarshal([]byte(update.CallbackQuery.Data), &data)
	ask := data.Ask
	i := 0
	for _, w := range product.FindItems(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10), ask) {
		file := tgbotapi.FileID("AgACAgIAAxkBAAIGomLrzy2KPw72xfdESzZ38rTCaBi7AALmwDEbmLVhSzsBlE2C-TvzAQADAgADcwADKQQ") // NGGYU picture
		if len(w.TitleFoto) >= 60 {
			file = tgbotapi.FileID(w.TitleFoto)
		}
		msg := tgbotapi.NewPhoto(update.CallbackQuery.Message.Chat.ID, file)
		msg.Caption = fmt.Sprintf("%s %s %s\n%s %s\n", ///showallphotos_%s\n\n\n\n/deleteitem_%s",
			w.Manufacture,
			w.Model,
			w.ItemID,
			w.Color,
			w.Comments)
		// w.UniqID[len(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)):],
		// w.UniqID[len(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)):])
		//msg.ReplyMarkp = tgbotapi.NewRemoveKeybard(true)
		ph, err := json.Marshal(CallBackButton{Command: "photos", ItemIDbyUser: w.UniqID[len(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)):]})
		if err != nil {
			ph = nil
		}
		del, err := json.Marshal(CallBackButton{Command: "deleteask", ItemIDbyUser: w.UniqID[len(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)):]})
		if err != nil {
			del = nil
		}
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("All Photos", string(ph)),
				tgbotapi.NewInlineKeyboardButtonData("Delete", string(del)),
			),
		)
		// send
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		i++
	}
	if i <= 0 {
		mess := "Looking for " + ask + "\nnot found"
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, mess)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// not used (prev version was a carusel) ABANDONED:
func FindItemCarusel(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	ask := inputMessage.Text
	// for fun:
	// t := fmt.Sprintf("looking for %s  ...", ask)
	// msg := tgbotapi.NewMessage(inputMessage.Chat.ID, t)
	// a, err := bot.Send(msg)
	// if err != nil {
	// 	log.Panic(err)
	// }
	//---------
	ww := product.FindItems(strconv.FormatInt(inputMessage.Chat.ID, 10), ask)

	// if not found
	if len(ww) == 0 {
		mess := "Looking for " + ask + "\nbut not found"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, mess)
		//msg.ReplyMarkp = tgbotapi.NewRemoveKeboard(true)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		return
	}

	sort.Slice(ww, func(i, j int) (less bool) {
		return ww[i].UniqID < ww[j].UniqID
	})

	wt := NextItemButtinCarusel{
		Task: "nextitem",
		// Chatid:    inputMessage.Chat.ID,
		// Messageid: a.MessageID,
		Word:  ask,
		Index: 1,
		//List:      ww,
	}
	//file := tgbotapi.FileID("AgACAgIAAxkDAAIFLhUvVhKqcrdTToDwPgQ3xvIUWxAAJ8vTEbDQ3wSt4fooC36sHLAQADAgADdwADKQQ")
	w := ww[0]
	mess := fmt.Sprintf("%d/%d\n", 1, len(ww))
	mess = mess + fmt.Sprintf("%s %s %s\n%s %s",
		w.Manufacture,
		w.Model,
		w.ItemID,
		w.Color,
		w.Comments)

	serializedData, e := json.Marshal(wt)
	if e != nil {
		log.Println(e)
	}
	log.Println(serializedData)

	file := tgbotapi.FileID("AgACAgIAAxkBAAIGAmLpQ12qYRKTL-4aT8f022NzKz_OAALevzEb-6pJS9Hg5i87VdjUAQADAgADeAADKQQ")
	if len(w.TitleFoto) >= 60 {
		file = tgbotapi.FileID(w.TitleFoto)
	}
	msg := tgbotapi.NewPhoto(inputMessage.Chat.ID, file)
	msg.Caption = mess
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next Page", string(serializedData)),
		),
	)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func FindNextItemCarusel(bot *tgbotapi.BotAPI, wt *NextItemButtinCarusel, update tgbotapi.Update) {
	ask := wt.Word
	ww := product.FindItems(strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10), ask)

	// if not found
	if len(ww) == 0 {

		mess := "Looking for " + ask + "\nnot found"
		msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, mess)
		//msg.ReplyMarkp = tgbotapi.NewRemoveKeboard(true)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		return
	}

	sort.Slice(ww, func(i, j int) (less bool) {
		return ww[i].Manufacture < ww[j].Manufacture
	})
	i := wt.Index

	//file := tgbotapi.FileID("AgACAgIAAxkDAAIFLhUvVhKqcrdTToDwPgQ3xvIUWxAAJ8vTEbDQ3wSt4fooC36sHLAQADAgADdwADKQQ")
	w := ww[i]
	wt.Index = i + 1

	mess := fmt.Sprintf("%d/%d\n", wt.Index, len(ww))
	mess = mess + fmt.Sprintf("%s %s %s\n%s %s",
		w.Manufacture,
		w.Model,
		w.ItemID,
		w.Color,
		w.Comments)

	if wt.Index == len(ww) {
		wt.Index = 0
	}

	serializedData, e := json.Marshal(wt)
	if e != nil {
		log.Println(e)
	}
	log.Println(serializedData)

	file := tgbotapi.FileID("AgACAgIAAxkBAAIGAmLpQ12qYRKTL-4aT8f022NzKz_OAALevzEb-6pJS9Hg5i87VdjUAQADAgADeAADKQQ")
	if len(w.TitleFoto) >= 60 {
		file = tgbotapi.FileID(w.TitleFoto)
	}
	// из примера:
	// msg:= tgbotapi.EditMessageMediaConfig{
	// 	BaseEdit: tgbotapi.BaseEdit{
	// 		MessageID: update.CallbackQuery.Message.MessageID,
	// 		ChatID: update.CallbackQuery.Message.Chat.ID, },
	// 	Media: tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(file)),
	// bot.Send(msg)

	rpl := tgbotapi.NewInlineKeyboardMarkup( //func NewEditMessageReplyMarkup(chatID int64, messageID int, replyMarkup InlineKeyboardMarkup) EditMessageReplyMarkupConfig
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next item", string(serializedData)), // только 64 bytes
		),
	)
	msg := tgbotapi.EditMessageMediaConfig{
		BaseEdit: tgbotapi.BaseEdit{
			MessageID:   update.CallbackQuery.Message.MessageID,
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			ReplyMarkup: &rpl},
		Media: tgbotapi.NewInputMediaPhoto(file),
	}
	//msg.Caption = mess
	//SendMediaGroup

	//msg.Caption = mess
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	// file := tgbotapi.FileID(w.TitleFoto)
	// msg := tgbotapi.NewPhoto(update.CallbackQuery.Message.Chat.ID, file)
	// msg.Caption = mess
	// msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("Next Page", string(serializedData)),
	// 	),
	// )
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }
}

// func NextItmCarusl(bot *tgboapi.BotAPI, update tgbotapi.Update, wt *NextItemButtinCarusel) {
// 	mess := "именен"
// 	serializedata, _ := json.Marsal(wt)
// 	btn := tgbotapi.NewInlineKebordMarkup(
// 		gotapi.NewIlineKeyboardRow(
// 		botapi.NewInlineKeyboardButtonData("Next Page", string(serializedData)),
// 		,
// 	)
// 	msgr := tgbotai.NewEditMessagTextAndMrkup(update.CallbackQuery.Message.Chat.ID, wt.Messageid, mess, btns)
// _, err = bot.end(msgr); err != nil {
// 	ogPanic(err)
//
// }
