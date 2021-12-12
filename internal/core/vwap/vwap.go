package vwap

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

const defaultMaxSize = 200

type CalculatorService struct {
	mu                sync.Mutex
	DataPointQueue    []DataPoint
	SumWeightedVolume map[string]decimal.Decimal
	SumVolume         map[string]decimal.Decimal
	VWAP              map[string]decimal.Decimal
	MaxSize           uint
}

func NewVwapCalculator(maxSize uint) *CalculatorService {
	return &CalculatorService{
		DataPointQueue:    []DataPoint{},
		MaxSize:           maxSize,
		SumWeightedVolume: make(map[string]decimal.Decimal),
		SumVolume:         make(map[string]decimal.Decimal),
		VWAP:              make(map[string]decimal.Decimal),
	}
}

func ProvideVwapCalculator() Calculator {
	return NewVwapCalculator(defaultMaxSize)
}

func (s *CalculatorService) Len() int {
	return len(s.DataPointQueue)
}

func (s *CalculatorService) Calculate(dataPoint DataPoint) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.DataPointQueue) == int(s.MaxSize) {
		d := s.DataPointQueue[0]
		s.DataPointQueue = s.DataPointQueue[1:]

		s.SumWeightedVolume[d.ProductID] = s.SumWeightedVolume[d.ProductID].Sub(d.Price.Mul(d.Volume))
		s.SumVolume[d.ProductID] = s.SumVolume[d.ProductID].Sub(d.Volume)

		if !s.SumVolume[d.ProductID].IsZero() {
			s.VWAP[d.ProductID] = s.SumWeightedVolume[d.ProductID].Div(s.SumVolume[d.ProductID])
		}
	}

	_, ok := s.VWAP[dataPoint.ProductID]
	switch ok {
	case true:
		s.SumWeightedVolume[dataPoint.ProductID] = s.SumWeightedVolume[dataPoint.ProductID].Add(dataPoint.Price.Mul(dataPoint.Volume))
		s.SumVolume[dataPoint.ProductID] = s.SumVolume[dataPoint.ProductID].Add(dataPoint.Volume)
		s.VWAP[dataPoint.ProductID] = s.SumWeightedVolume[dataPoint.ProductID].Div(s.SumVolume[dataPoint.ProductID])
	default:
		firstVW := dataPoint.Price.Mul(dataPoint.Volume)

		s.SumWeightedVolume[dataPoint.ProductID] = firstVW
		s.SumVolume[dataPoint.ProductID] = dataPoint.Volume
		s.VWAP[dataPoint.ProductID] = firstVW.Div(dataPoint.Volume)
	}

	s.DataPointQueue = append(s.DataPointQueue, dataPoint)

	fmt.Println(s.VWAP)
}
