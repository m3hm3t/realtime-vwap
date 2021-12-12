package vwap

type Calculator interface {
	Calculate(data DataPoint)
	Len() int
}
