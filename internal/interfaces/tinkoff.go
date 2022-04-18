package interfaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/WidowGenerator/tinkoff_dollars_bot/internal/db/enums"
)

const clustersUrl string = "https://api.tinkoff.ru/geo/withdraw/clusters"

var ErrBadStatusCode = errors.New("interfaces/tinkoff: Bad status code")

func NewTinkkoffAtmAPI(tinkoffATMBody *TinkoffATMBody, City *enums.City, Amount *uint) TinkoffATM {
	return TinkoffATM{
		TinkoffATMBody: tinkoffATMBody,
		City:           City,
		Amount:         Amount,
	}
}

func (tinkoffAtm *TinkoffATM) GetATMs() error {
	data := TinkoffATMResponse{}

	json_data, err := json.Marshal(tinkoffAtm.TinkoffATMBody)

	if err != nil {
		return err
	}

	resp, err := http.Post(clustersUrl, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return ErrBadStatusCode
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(body, &data)

	tinkoffAtm.TinkoffATMResponse = &data

	return nil
}

func (tinkoffAtm *TinkoffATM) GetMessage() (string, error) {
	message := fmt.Sprintf(
		`
Info by atms in %v
		`,
		tinkoffAtm.City,
	)

	moreBankInfo := ""

	for i := 0; i < len(tinkoffAtm.TinkoffATMResponse.Payload.Clusters); i++ {
		cluster := tinkoffAtm.TinkoffATMResponse.Payload.Clusters[i]

		for j := 0; j < len(cluster.Points); j++ {
			point := cluster.Points[j]

			k := 0
			for k = 0; k < len(point.Limits); k++ {
				if point.Limits[k].Currency == tinkoffAtm.TinkoffATMBody.Filters.Currencies[0] {
					break
				}
			}

			if point.Limits[k].Amount < *tinkoffAtm.Amount {
				continue
			}

			moreBankInfo += fmt.Sprintf(
				`
				Address: %v
				ID: %v
				Contacts: %v
				Atm: %v
				Amount: %v (%v)
				`,
				point.Address,
				point.Id,
				point.Phone[0],
				point.Brand.Name,
				point.Limits[k].Amount,
				point.Limits[k].Currency,
			)
		}
	}

	if len(moreBankInfo) == 0 {
		return message + "There are no suitable ATMs nearby", nil
	}

	return message + moreBankInfo, nil
}
