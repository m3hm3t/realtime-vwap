package vwap

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

const defaultMaxSize = 200

type CalculatorService struct {
	mu                sync.Mutex
	DataPoints        []DataPoint
	SumVolumeWeighted map[string]decimal.Decimal
	SumVolume         map[string]decimal.Decimal
	VWAP              map[string]decimal.Decimal

	MaxSize uint
}

func NewVwapCalculator(maxSize uint) *CalculatorService {
	return &CalculatorService{
		DataPoints:        []DataPoint{},
		MaxSize:           maxSize,
		SumVolumeWeighted: make(map[string]decimal.Decimal),
		SumVolume:         make(map[string]decimal.Decimal),
		VWAP:              make(map[string]decimal.Decimal),
	}
}

func ProvideVwapCalculator() Calculator {
	return NewVwapCalculator(defaultMaxSize)
}

func (s *CalculatorService) Len() int {
	return len(s.DataPoints)
}

func (s *CalculatorService) Calculate(dataPoint DataPoint) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.DataPoints) == int(s.MaxSize) {
		d := s.DataPoints[0]
		s.DataPoints = s.DataPoints[1:]

		s.SumVolumeWeighted[d.ProductID] = s.SumVolumeWeighted[d.ProductID].Sub(d.Price.Mul(d.Volume))
		s.SumVolume[d.ProductID] = s.SumVolume[d.ProductID].Sub(d.Volume)
		if !s.SumVolume[d.ProductID].IsZero() {
			s.VWAP[d.ProductID] = s.SumVolumeWeighted[d.ProductID].Div(s.SumVolume[d.ProductID])
		}
	}

	if _, ok := s.VWAP[dataPoint.ProductID]; ok {
		s.SumVolumeWeighted[dataPoint.ProductID] = s.SumVolumeWeighted[dataPoint.ProductID].Add(dataPoint.Price.Mul(dataPoint.Volume))
		s.SumVolume[dataPoint.ProductID] = s.SumVolume[dataPoint.ProductID].Add(dataPoint.Volume)
		s.VWAP[dataPoint.ProductID] = s.SumVolumeWeighted[dataPoint.ProductID].Div(s.SumVolume[dataPoint.ProductID])
	} else {
		initialVW := dataPoint.Price.Mul(dataPoint.Volume)

		s.SumVolumeWeighted[dataPoint.ProductID] = initialVW
		s.SumVolume[dataPoint.ProductID] = dataPoint.Volume
		s.VWAP[dataPoint.ProductID] = initialVW.Div(dataPoint.Volume)
	}

	s.DataPoints = append(s.DataPoints, dataPoint)

	fmt.Println(s.VWAP)
}
