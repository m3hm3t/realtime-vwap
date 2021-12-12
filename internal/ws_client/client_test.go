package ws_client_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/mock"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/subscriber"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestWebsocketSubscribe(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		Desc         string
		TradingPairs []string
		ExpectError  bool
	}{
		{
			Desc:         "ValidPair",
			TradingPairs: []string{"ETH-BTC"},
			ExpectError:  false,
		},
		{
			Desc:         "InvalidPair",
			TradingPairs: []string{"xxx-BTC"},
			ExpectError:  true,
		},
	}

	mockPairHandler := mock.NewMockHandler(ctrl)
	pairSubs := subscriber.NewPairSubscriber()
	pairWsClient := ws_client.NewClient(pairSubs, mockPairHandler)

	for _, tt := range tests {
		tt := tt

		t.Run(tt.Desc, func(t *testing.T) {
			t.Parallel()

			var wg sync.WaitGroup

			if tt.ExpectError {
				mockPairHandler.EXPECT().Handle(gomock.Any()).Times(0)
			} else {
				wg.Add(1)
				mockPairHandler.EXPECT().Handle(gomock.Any()).
					Do(func(_ interface{}) {
						wg.Done()
					}).MinTimes(1)
			}

			err := pairWsClient.Start(ctx, tt.TradingPairs)
			wg.Wait()

			if tt.ExpectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
