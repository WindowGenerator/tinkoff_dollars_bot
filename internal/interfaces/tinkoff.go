package interfaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const clustersUrl string = "https://api.tinkoff.ru/geo/withdraw/clusters"

var testValues = TinkoffATMBody{
	Bounds: &Bounds{
		BottomLeft: GeographicCoodrinates{
			Lat: 56.63827841259033,
			Lng: 60.33014140617972,
		},
		TopRight: GeographicCoodrinates{
			Lat: 57.005757763347255,
			Lng: 60.9948142577422,
		},
	},
	Filters: &Filters{
		ShowUnavailable: true,
		Currencies:      []string{"USD"},
	},
	Zoom: 11,
}

var ErrBadStatusCode = errors.New("interfaces/tinkoff: Bad status code")

func NewTinkkoffAtmAPI(bounds *Bounds, filters *Filters, zoom uint32) TinkoffATM {
	return TinkoffATM{
		TinkoffATMBody: &TinkoffATMBody{
			Bounds:  bounds,
			Filters: filters,
			Zoom:    zoom,
		},
	}
}

func (tinkoffAtm *TinkoffATM) GetATMs() error {
	data := TinkoffATMResponse{}

	json_data, err := json.Marshal(testValues)

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
	message := `Информация по банкоматам в городе Екатеринбург\n\n`
	if len(tinkoffAtm.TinkoffATMResponse.Payload.Clusters) == 0 {
		return message + "Рядом нет подходящих банкоматов", nil
	}

	return message, nil
}
