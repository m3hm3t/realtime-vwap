package subscriber

import (
	"encoding/json"
	"fmt"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/dto"
	ws "golang.org/x/net/websocket"
)

type PairSubscriber struct {
}

func NewPairSubscriber() Subscriber {
	return &PairSubscriber{}
}

func ProvidePairSubscriber() Subscriber {
	return NewPairSubscriber()
}

func (s *PairSubscriber) Subscribe(conn *ws.Conn, pairs []string) error {

	subscription := dto.Request{
		Type:       dto.RequestTypeSubscribe,
		ProductIDs: pairs,
		Channels: []dto.Channel{
			{Name: dto.ChannelTypeMatches},
		},
	}

	payload, err := json.Marshal(subscription)
	if err != nil {
		return fmt.Errorf("failed to marshal subscription: %w", err)
	}

	err = ws.Message.Send(conn, payload)
	if err != nil {
		return fmt.Errorf("failed to send subscription: %w", err)
	}

	var subscriptionResponse dto.Response
	err = ws.JSON.Receive(conn, &subscriptionResponse)
	if err != nil {
		return fmt.Errorf("failed to receive subscription response: %w", err)
	}

	if subscriptionResponse.Type == "error" {
		return fmt.Errorf("failed to subscribe: %s", subscriptionResponse.Message)
	}

	return nil
}
