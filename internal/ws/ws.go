package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/m3hm3t/realtime-vwap/internal/ws/dto"
	"log"
	"strings"
)

const DefaultURL = "wss://ws-feed.exchange.coinbase.com"
const defaultTradingPairs = "BTC-USD,ETH-USD,ETH-BTC"

func RegisterHandlers(receiver chan dto.Response) {

}

func Run(ctx context.Context) error {

	conn, _, err := websocket.DefaultDialer.Dial(DefaultURL, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	subscription := dto.Request{
		Type:       dto.RequestTypeSubscribe,
		ProductIDs: strings.Split(defaultTradingPairs, ","),
		Channels: []dto.Channel{
			{Name: "matches"},
		},
	}

	payload, err := json.Marshal(subscription)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription: %w", err)
	}

	err = conn.WriteJSON(payload)
	if err != nil {
		return fmt.Errorf("failed to send subscription: %w", err)
	}

	var subscriptionResponse dto.Response
	err = conn.ReadJSON(&subscriptionResponse)
	if err != nil {
		return fmt.Errorf("failed to receive subscription response: %w", err)
	}

	receiver := make(chan dto.Response)

	go func() {
		for {
			select {
			case <-ctx.Done():
				err := conn.Close()
				if err != nil {
					log.Printf("failed closing ws connection: %s", err)
				}
			default:
				var response dto.Response

				err := conn.ReadJSON(&response)
				if err != nil {
					log.Printf("failed receiving response: %s", err)

					break
				}

				receiver <- response
			}
		}
	}()

	return nil
}
