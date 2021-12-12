package ws_client

import (
	"context"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/dto"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/handler"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/subscriber"
	ws "golang.org/x/net/websocket"
	"log"
)

const DefaultURL = "wss://ws-feed.exchange.coinbase.com"

type Client struct {
	handler    handler.Handler
	subscriber subscriber.Subscriber
}

func NewClient(subscriber subscriber.Subscriber, handler handler.Handler) Client {
	return Client{subscriber: subscriber, handler: handler}
}

func ProvideWsClient(subscriber subscriber.Subscriber, handler handler.Handler) Client {
	return NewClient(subscriber, handler)
}

func (client *Client) Start(ctx context.Context, pairs []string) error {

	conn, err := ws.Dial(DefaultURL, "", "http://localhost/")
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Printf("websocket connected to: %s", DefaultURL)

	if err := client.subscriber.Subscribe(conn, pairs); err != nil {
		return err
	}

	receiver := make(chan dto.Response)

	go func() {
		for {
			select {
			case <-ctx.Done():
				err := conn.Close()
				if err != nil {
					log.Printf("failed closing ws_client connection: %s", err)
				}
			default:
				var response dto.Response

				err := ws.JSON.Receive(conn, &response)
				if err != nil {
					log.Printf("failed receiving message: %s", err)

					break
				}

				receiver <- response
			}
		}
	}()

	client.handler.Handle(receiver)

	return nil
}
