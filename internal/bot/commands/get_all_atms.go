package commands

import (
	"flag"
	"log"
	"regexp"
	"strings"

	db "github.com/WidowGenerator/tinkoff_dollars_bot/internal/db"
	enums "github.com/WidowGenerator/tinkoff_dollars_bot/internal/db/enums"
	interfaces "github.com/WidowGenerator/tinkoff_dollars_bot/internal/interfaces"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var spaceFindExpression = regexp.MustCompile("\\s+")

type InputAtmInfo struct {
	City            *enums.City
	Currencies      *[]string
	ShowUnavailable *bool
}

// Help display the help for the bot
func GetAllAtms(chatID int64, bot *tgbotapi.BotAPI, arguments string) error {

	msg := tgbotapi.NewMessage(chatID, "")
	tinkoffATMBody, err := parseGetAllAtmsArgs(arguments)

	if err != nil {
		msg.Text = helpMessage
		_, err = bot.Send(msg)

		return nil
	}

	tinkoffATM := interfaces.NewTinkkoffAtmAPI(tinkoffATMBody.Bounds, tinkoffATMBody.Filters, tinkoffATMBody.Zoom)

	err = tinkoffATM.GetATMs()

	if err != nil {
		log.Fatal(err)
	}

	text, err := tinkoffATM.GetMessage()

	if err != nil {
		log.Fatal(err)
	}

	msg.Text = text
	_, err = bot.Send(msg)

	return err
}

func parseGetAllAtmsArgs(args string) (interfaces.TinkoffATMBody, error) {
	splitedArgs := strings.Split(args, " ")
	tinkoffATMBody := interfaces.TinkoffATMBody{}

	fs, inputAtmInfo := getAtmParser()

	err := fs.Parse(splitedArgs)

	if err != nil {
		return tinkoffATMBody, err
	}

	boundsWithZoom := db.Cities2GS[*inputAtmInfo.City]

	tinkoffATMBody.Bounds = boundsWithZoom.Bounds
	tinkoffATMBody.Filters = &interfaces.Filters{
		ShowUnavailable: *inputAtmInfo.ShowUnavailable,
		Currencies:      *inputAtmInfo.Currencies,
	}
	tinkoffATMBody.Zoom = boundsWithZoom.Zoom

	return tinkoffATMBody, nil
}

func getAtmParser() (*flag.FlagSet, *InputAtmInfo) {
	fs := flag.NewFlagSet("GetAllAtms", flag.ContinueOnError)
	inputAtmInfo := InputAtmInfo{}

	inputAtmInfo.City = enums.CityMoscow.Ptr()
	fs.Func("city", "Please print the city, which you want to see", func(s string) error {
		city, err := enums.ParseCity(s)

		if err != nil {
			return err
		}

		inputAtmInfo.City = city.Ptr()

		return nil
	})
	inputAtmInfo.ShowUnavailable = fs.Bool("show-unavailable", false, "If you want to see unavailable atms")
	fs.Func("currencies", "aaaaaaaa", func(s string) error {
		splitedCurrencies := spaceFindExpression.Split(strings.TrimSpace(s), -1)
		convertedCurrencies := make([]string, len(splitedCurrencies))

		for i := 0; i < len(splitedCurrencies); i += 1 {
			_, err := enums.ParseCurrency(splitedCurrencies[i])
			if err != nil {
				return err
			}
			convertedCurrencies[i] = splitedCurrencies[i]
		}

		inputAtmInfo.Currencies = &convertedCurrencies

		return nil
	})

	return fs, &inputAtmInfo
}
