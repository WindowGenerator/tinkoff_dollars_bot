package interfaces

import "github.com/WidowGenerator/tinkoff_dollars_bot/internal/db/enums"

type Bounds struct {
	BottomLeft GeographicCoodrinates `json:"bottomLeft"`
	TopRight   GeographicCoodrinates `json:"topRight"`
}

type GeographicCoodrinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Filters struct {
	Banks           []string `json:"banks"`
	Currencies      []string `json:"currencies"`
	ShowUnavailable bool     `json:"showUnavailable"`
}

type Brand struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LogoFile    string `json:"logoFile"`
	RoundedLogo bool   `json:"roundedLogo"`
}

type Limit struct {
	Currency      string `json:"currency"`
	Max           uint   `json:"max"`
	Denominations []uint `json:"denominations"`
	Amount        uint   `json:"amount"`
}

type WorkPeriods struct {
	OpenDay   uint   `json:"openDay"`
	OpenTime  string `json:"openTime"`
	CloseDay  uint   `json:"closeDay"`
	CloseTime string `json:"closeTime"`
}

type AtmStatuses struct {
	CriticalFailure       bool `json:"criticalFailure"`
	QrOperational         bool `json:"qrOperational"`
	NfcOperational        bool `json:"nfcOperational"`
	CardReaderOperational bool `json:"cardReaderOperational"`
	CashInAvailable       bool `json:"cashInAvailable"`
}

type AtmInfo struct {
	Available  bool        `json:"available"`
	IsTerminal bool        `json:"isTerminal"`
	Statuses   AtmStatuses `json:"statuses"`
}

type Point struct {
	Id           string                `json:"id"`
	Brand        Brand                 `json:"brand"`
	PointType    string                `json:"pointType"`
	Location     GeographicCoodrinates `json:"location"`
	Address      string                `json:"address"`
	Phone        []string              `json:"phone"`
	Limits       []Limit               `json:"limits"`
	WorkPeriods  WorkPeriods           `json:"workPeriods"`
	InstallPlace string                `json:"installPlace"`
	AtmInfo      AtmInfo               `json:"atmInfo"`
}

type Cluster struct {
	Id     string                `json:"id"`
	Hash   string                `json:"hash"`
	Bounds Bounds                `json:"bounds"`
	Center GeographicCoodrinates `json:"center"`
	Points []Point               `json:"points"`
}

type Payload struct {
	Hash     string    `json:"hash"`
	Zoom     uint32    `json:"zoom"`
	Bounds   Bounds    `json:"bounds"`
	Clusters []Cluster `json:"clusters"`
}

type TinkoffATMBody struct {
	Bounds  *Bounds  `json:"bounds"`
	Filters *Filters `json:"filters"`
	Zoom    uint32   `json:"zoom"`
}

type TinkoffATMResponse struct {
	TrackingId string  `json:"trackingId"`
	Payload    Payload `json:"payload"`
	Time       string  `json:"time"`
	Status     string  `json:"status"`
}

type TinkoffATM struct {
	City               *enums.City
	Amount             *uint
	TinkoffATMBody     *TinkoffATMBody
	TinkoffATMResponse *TinkoffATMResponse
}
