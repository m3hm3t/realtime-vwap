package handler

import "github.com/m3hm3t/realtime-vwap/internal/ws_client/dto"

type Handler interface {
	Handle(receiver chan dto.Response)
}
