package vwap_test

import (
	"github.com/m3hm3t/realtime-vwap/internal/core/vwap"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestList(t *testing.T) {
	t.Parallel()

	vwapCalculator := vwap.NewVwapCalculator(1)

	first := vwap.DataPoint{Price: decimal.NewFromInt(1), Volume: decimal.NewFromInt(1)}

	second := vwap.DataPoint{Price: decimal.NewFromInt(2), Volume: decimal.NewFromInt(2)}

	third := vwap.DataPoint{Price: decimal.NewFromInt(3), Volume: decimal.NewFromInt(3)}

	vwapCalculator.Calculate(first)
	assert.Equal(t, 1, vwapCalculator.Len())
	assert.Equal(t, first, vwapCalculator.DataPoints[0])

	vwapCalculator.Calculate(second)
	assert.Equal(t, 1, vwapCalculator.Len())
	assert.Equal(t, second, vwapCalculator.DataPoints[0])

	vwapCalculator.Calculate(third)
	assert.Equal(t, 1, vwapCalculator.Len())
	assert.Equal(t, third, vwapCalculator.DataPoints[0])
}

func TestListConcurrentPush(t *testing.T) {
	t.Parallel()

	vwapCalculator := vwap.NewVwapCalculator(2)

	first := vwap.DataPoint{Price: decimal.NewFromInt(1), Volume: decimal.NewFromInt(1)}

	second := vwap.DataPoint{Price: decimal.NewFromInt(2), Volume: decimal.NewFromInt(2)}

	third := vwap.DataPoint{Price: decimal.NewFromInt(3), Volume: decimal.NewFromInt(3)}

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		vwapCalculator.Calculate(first)
		wg.Done()
	}()

	go func() {
		vwapCalculator.Calculate(second)
		wg.Done()
	}()

	go func() {
		vwapCalculator.Calculate(third)
		wg.Done()
	}()

	wg.Wait()

	assert.Equal(t, vwapCalculator.Len(), 2)
}

func TestVWAP(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name       string
		DataPoints []vwap.DataPoint
		WantVwap   map[string]decimal.Decimal
		MaxSize    uint
	}{
		{
			Name:       "EmptyDataPoints",
			DataPoints: []vwap.DataPoint{},
			WantVwap: map[string]decimal.Decimal{
				"BTC-USD": decimal.Zero,
				"ETH-USD": decimal.Zero,
			},
		},
		{
			Name: "FullDataPoints1",
			DataPoints: []vwap.DataPoint{
				{Price: decimal.NewFromInt(10), Volume: decimal.NewFromInt(10), ProductID: "BTC-USD"},
				{Price: decimal.NewFromInt(10), Volume: decimal.NewFromInt(10), ProductID: "BTC-USD"},
				{Price: decimal.NewFromInt(31), Volume: decimal.NewFromInt(30), ProductID: "ETH-USD"},
				{Price: decimal.NewFromInt(21), Volume: decimal.NewFromInt(20), ProductID: "BTC-USD"},
				{Price: decimal.NewFromInt(41), Volume: decimal.NewFromInt(33), ProductID: "ETH-USD"},
			},
			MaxSize: 4,
			WantVwap: map[string]decimal.Decimal{
				"BTC-USD": decimal.RequireFromString("17.3333333333333333"),
				"ETH-USD": decimal.RequireFromString("36.2380952380952381"),
			},
		},
		{
			Name: "FullDataPoints2",
			DataPoints: []vwap.DataPoint{
				{Price: decimal.NewFromInt(10), Volume: decimal.RequireFromString("10.1"), ProductID: "BTC-USD"},
				{Price: decimal.NewFromInt(10), Volume: decimal.RequireFromString("10.1"), ProductID: "BTC-USD"},
			},
			WantVwap: map[string]decimal.Decimal{
				"BTC-USD": decimal.RequireFromString("10"),
			},
			MaxSize: 4,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			vwapCalculator := vwap.NewVwapCalculator(tt.MaxSize)

			for _, d := range tt.DataPoints {
				vwapCalculator.Calculate(d)
			}

			for k := range tt.WantVwap {
				assert.Equal(t, tt.WantVwap[k].String(), vwapCalculator.VWAP[k].String())
			}
		})
	}
}
