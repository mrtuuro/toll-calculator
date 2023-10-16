package main

import (
	"fmt"
	"toll-calculator/types"
)

const basePrice = 3.15

type Aggregator interface {
	AggregateDistance(types.Distance) error
	CalculateInvoice(int) (*types.Invoice, error)
}

type Storer interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (ia *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("Processing and inserting distance to the storage. ", distance)
	return ia.store.Insert(distance)
}

func (ia *InvoiceAggregator) CalculateInvoice(obuID int) (*types.Invoice, error) {
	dist, err := ia.store.Get(obuID)
	if err != nil {
		return nil, err
	}
	inv := &types.Invoice{
		OBUID:         obuID,
		TotalDistance: dist,
		TotalAmount:   basePrice * dist,
	}
	return inv, nil
}
