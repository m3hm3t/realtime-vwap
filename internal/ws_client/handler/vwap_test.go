package handler_test

import (
	"github.com/golang/mock/gomock"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/dto"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/handler"
	"github.com/m3hm3t/realtime-vwap/internal/ws_client/handler/mock"
	"sync"
	"testing"
)

func TestShouldHandleResponse(t *testing.T) {
	t.Parallel()

	// GIVEN
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var wg sync.WaitGroup

	numberOfTests := [5]struct{}{}

	stubResponse := dto.Response{
		Type:      "mock",
		Channels:  nil,
		Message:   "mock",
		Size:      "2.2",
		Price:     "1.1",
		ProductID: "mock",
	}
	stubModel := stubResponse.BuildModel()

	mockVwapCalculator := mock.NewMockCalculator(ctrl)
	mockVwapCalculator.EXPECT().Calculate(stubModel).
		Do(func(_ interface{}) {
			wg.Done()
		}).Times(len(numberOfTests))

	vwapHandler := handler.NewVwapHandler(mockVwapCalculator)

	receiver := make(chan dto.Response)

	// WHEN
	go vwapHandler.Handle(receiver)

	for range numberOfTests {
		wg.Add(1)
		receiver <- stubResponse
	}

	// THEN
	wg.Wait()
}
