package main

import (
	"log"
)

//type DistanceCalculator struct {
//	consumer DataConsumer
//}

const kafkaTopic = "obu-data"

// Transport (HTTP, gRPC, Kafka) -> attach business logic to this transport

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
