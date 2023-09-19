package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"toll-calculator/types"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleware{
		next: next,
	}
}

func (mw *LogMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": fmt.Sprintf("%.2f", dist),
			"long": data.Long,
			"lat":  data.Lat,
		}).Info("calculate distance")
	}(time.Now())
	dist, err = mw.next.CalculateDistance(data)
	return
}
