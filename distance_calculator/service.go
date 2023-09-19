package main

import (
	"math"
	"toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	prevPoint []float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{}
}

func (cs *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0
	if len(cs.prevPoint) > 0 {
		distance = calculateDistance(cs.prevPoint[0], cs.prevPoint[1], data.Lat, data.Long)
	}
	cs.prevPoint = []float64{data.Long, data.Lat}
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
