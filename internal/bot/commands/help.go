package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const helpMessage string = `
Help:

- help
`

// Help display the help for the bot
func Help(chatID int64, bot *tgbotapi.BotAPI, arguments string) error {

	msg := tgbotapi.NewMessage(chatID, "this is the help")

	_, err := bot.Send(msg)

	return err
}
