package vwap

import "github.com/shopspring/decimal"

type DataPoint struct {
	Price     decimal.Decimal
	Volume    decimal.Decimal
	ProductID string
}
