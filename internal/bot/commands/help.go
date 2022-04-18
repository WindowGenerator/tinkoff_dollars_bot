package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const helpMessage string = `
Help:

- help: show help message
- start: start discuss wuth bot
- get_atms: get atms by filter (if you want to know more about this command type help get_atms)
`

// Help display the help for the bot
func Help(chatID int64, bot *tgbotapi.BotAPI, arguments string) error {

	msg := tgbotapi.NewMessage(chatID, helpMessage)

	if arguments == "get_atms" {
		msg.Text = GetAtmsHelpMessage
	}

	_, err := bot.Send(msg)

	return err
}
