package main

import (
	"log"

	bot "github.com/WidowGenerator/tinkoff_dollars_bot/internal/bot"
	helper "github.com/WidowGenerator/tinkoff_dollars_bot/internal/helpers"
)

// main function start the application.
// First load the file .env with the environment variables
// and ask for the token to call the main bot initializer
func main() {
	if err := helper.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	token := helper.GetBotToken()

	log.Fatal(bot.StartBot(token))
}
