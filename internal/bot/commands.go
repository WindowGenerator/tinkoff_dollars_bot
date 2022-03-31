package bot

import (
	"log"

	command "github.com/WidowGenerator/tinkoff_dollars_bot/internal/bot/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var commandList = map[string]interface{}{
	"help":                command.Help,
	"start":               command.Start,
	"get_all_atms":        command.GetAllAtms,
	"setup_notifications": command.SetupNotofications,
}

func cmdHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {

	if update.Message.IsCommand() {

		commandReq := update.Message.Command()
		commandArgs := update.Message.CommandArguments()

		if _, ok := commandList[commandReq]; !ok {
			// Make a default action ?
			return nil
		}

		err := commandList[commandReq].(func(int64, *tgbotapi.BotAPI, string) error)(update.Message.Chat.ID, bot, commandArgs)

		if err != nil {
			log.Fatal(err)
		}

	}
	return nil
}

func commandHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	return cmdHandler(update, bot)
}
