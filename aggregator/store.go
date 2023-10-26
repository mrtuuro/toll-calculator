package main

import (
	"fmt"
	"toll-calculator/types"
)

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (m *MemoryStore) Insert(distance types.Distance) error {
	m.data[distance.OBUID] += distance.Value
	return nil
}

func (m *MemoryStore) Get(id int) (float64, error) {
	distance, ok := m.data[id]
	fmt.Println("---||", m.data)
	if !ok {
		return 0.0, fmt.Errorf("could not find distance for obu id %d", id)
	}
	return distance, nil
}
