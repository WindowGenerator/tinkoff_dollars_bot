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
	Currency        *enums.Currency
	Bank            *enums.Bank
	Amount          *uint
	ShowUnavailable *bool
}

var GetAtmsHelpMessage = ``

// Help display the help for the bot
func GetAtms(chatID int64, bot *tgbotapi.BotAPI, arguments string) error {

	msg := tgbotapi.NewMessage(chatID, "")
	tinkoffATMBody, inputAtmInfo, err := parseGetAtmsArgs(arguments)

	if err != nil {
		msg.Text = helpMessage
		_, err = bot.Send(msg)

		return nil
	}

	tinkoffATM := interfaces.NewTinkkoffAtmAPI(tinkoffATMBody, inputAtmInfo.City, inputAtmInfo.Amount)

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

func parseGetAtmsArgs(args string) (*interfaces.TinkoffATMBody, *InputAtmInfo, error) {
	splitedArgs := spaceFindExpression.Split(strings.TrimSpace(args), -1)

	fs, inputAtmInfo := getAtmParser()

	err := fs.Parse(splitedArgs)

	if err != nil {
		return &interfaces.TinkoffATMBody{}, inputAtmInfo, err
	}

	boundsWithZoom := db.Cities2GS[*inputAtmInfo.City]

	return &interfaces.TinkoffATMBody{
		Bounds: boundsWithZoom.Bounds,
		Filters: &interfaces.Filters{
			ShowUnavailable: *inputAtmInfo.ShowUnavailable,
			Currencies:      []string{inputAtmInfo.Currency.String()},
			Banks:           *getBanks(inputAtmInfo.Bank.Ptr()),
		},
		Zoom: boundsWithZoom.Zoom,
	}, inputAtmInfo, nil
}

func getAtmParser() (*flag.FlagSet, *InputAtmInfo) {
	fs := flag.NewFlagSet("GetAtms", flag.ContinueOnError)
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
	inputAtmInfo.Currency = enums.CurrencyUSD.Ptr()
	fs.Func("currency", "Specify specific currency", func(s string) error {
		cur, err := enums.ParseCurrency(s)
		if err != nil {
			return err
		}

		inputAtmInfo.Currency = &cur

		return nil
	})
	inputAtmInfo.Bank = enums.BankAll.Ptr()
	fs.Func("bank", "Specify a specific city", func(s string) error {
		bank, err := enums.ParseBank(s)
		if err != nil {
			return err
		}

		inputAtmInfo.Bank = &bank

		return nil
	})

	inputAtmInfo.Amount = fs.Uint("amount", 100, "text")

	return fs, &inputAtmInfo
}

func getBanks(bank *enums.Bank) *[]string {
	if *bank == enums.BankAll {
		return &db.ALL_BANKS
	} else {
		return &[]string{db.Bank2BankId[*bank]}
	}
}
