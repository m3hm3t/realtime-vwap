//go:build wireinject
// +build wireinject

package wired

import (
	"github.com/google/wire"
	"github.com/m3hm3t/realtime-vwap/internal/core/vwap"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/handler"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/subscriber"
)

var vwapCalculatorSet = wire.NewSet(
	subscriber.ProvidePairSubscriber,
	handler.ProvideVwapHandler,
	vwap.ProvideVwapCalculator,
)

func InitializeWsClient() ws_client.Client {

	wire.Build(ws_client.ProvideWsClient, vwapCalculatorSet)

	return ws_client.Client{}
}
