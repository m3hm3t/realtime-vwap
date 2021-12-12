package handler

import (
	"github.com/m3hm3t/realtime-vwap/internal/core/vwap"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/dto"
)

type VwapHandler struct {
	calculator vwap.Calculator
}

func NewVwapHandler(calculator vwap.Calculator) Handler {
	return &VwapHandler{calculator: calculator}
}

func ProvideVwapHandler(calculator vwap.Calculator) Handler {
	return NewVwapHandler(calculator)
}

func (h *VwapHandler) Handle(receiver chan dto.Response) {

	for data := range receiver {
		if data.Price == "" {
			continue
		}

		h.calculator.Calculate(data.BuildModel())
	}
}
