package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Help display the help for the bot
func Start(chatID int64, bot *tgbotapi.BotAPI, arguments string) error {

	msg := tgbotapi.NewMessage(chatID, helpMessage)
	_, err := bot.Send(msg)

	return err
}
