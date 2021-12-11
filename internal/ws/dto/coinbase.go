package dto

import (
	"github.com/m3hm3t/realtime-vwap/internal/core/vwap"
	"github.com/shopspring/decimal"
	"log"
)

type RequestType string

const (
	RequestTypeSubscribe     RequestType = "subscribe"
	RequestTypeUnsubscribe   RequestType = "unsubscribe"
	RequestTypeSubscriptions RequestType = "subscriptions"
	RequestTypeError         RequestType = "error"
)

type ChannelType string

const (
	ChannelTypeLevel2    ChannelType = "level2"
	ChannelTypeHeartBeat ChannelType = "heartbeat"
	ChannelTypeTicker    ChannelType = "ticker"
	ChannelTypeMatches   ChannelType = "matches"
)

type Channel struct {
	Name       ChannelType
	ProductIDs []string
}

type Request struct {
	Type       RequestType `json:"type"`
	ProductIDs []string    `json:"product_ids"`
	Channels   []Channel   `json:"channels"`
}

type Response struct {
	Type      string    `json:"type"`
	Channels  []Channel `json:"channels"`
	Message   string    `json:"message,omitempty"`
	Size      string    `json:"size"`
	Price     string    `json:"price"`
	ProductID string    `json:"product_id"`
}

func (r Response) BuildModel() vwap.DataPoint {
	decimalPrice, err := decimal.NewFromString(r.Price)
	if err != nil {
		log.Printf("decimalPrice %s: %s", r.Price, err.Error())
	}

	decimalSize, err := decimal.NewFromString(r.Size)
	if err != nil {
		log.Printf("decimalPrice %s: %s", r.Price, err.Error())
	}

	return vwap.DataPoint{
		Price:     decimalPrice,
		Volume:    decimalSize,
		ProductID: r.ProductID,
	}
}
