package handler

import (
	"github.com/m3hm3t/realtime-vwap/internal/core/vwap"
	"github.com/m3hm3t/realtime-vwap/internal/ws/dto"
)

type VwapHandler struct {
	calculator vwap.Calculator
	receiver   chan dto.Response
}

func NewVwapHandler(receiver chan dto.Response) VwapHandler {
	return VwapHandler{receiver: receiver}
}

func (h *VwapHandler) Handle() {

	for data := range h.receiver {
		if data.Price == "" {
			continue
		}

		h.calculator.Calculate(data.BuildModel())
	}
}
