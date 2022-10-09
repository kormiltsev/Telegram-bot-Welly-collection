package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// just for future
func Quiz(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	poll := tgbotapi.SendPollConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: update.Message.Chat.ID,
		},
		Question:        "Answer to the Ultimate Question of Life, the Universe, and Everything",
		Type:            "quiz",
		Options:         []string{"41", "42", "32", "1"},
		CorrectOptionID: 1,
	}
	bot.Send(poll)
}
