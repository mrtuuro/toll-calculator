package main

import (
	"github.com/sirupsen/logrus"
	"time"
	"toll-calculator/types"
)

type LogMiddleware struct {
	next Aggregator
}

func (m *LogMiddleware) DistanceSum(i int) (float64, error) {
	//TODO implement me
	return 0, nil
}

func NewLogMiddleware(next Aggregator) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(start),
			"error": err,
		}).Info("Aggregate distance")
	}(time.Now())
	err = m.next.AggregateDistance(distance)
	return
}

func (m *LogMiddleware) CalculateInvoice(obuID int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if inv != nil {
			distance = inv.TotalDistance
			amount = inv.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"took":     time.Since(start),
			"error":    err,
			"obuID":    obuID,
			"amount":   amount,
			"distance": distance,
		}).Info("Aggregate distance")
	}(time.Now())
	inv, err = m.next.CalculateInvoice(obuID)
	return
}
