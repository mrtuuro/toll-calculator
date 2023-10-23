package main

import (
	"log"
	"toll-calculator/aggregator/client"
)

//type DistanceCalculator struct {
//	consumer DataConsumer
//}

const (
	kafkaTopic         = "obu-data"
	aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"
)

// Transport (HTTP, gRPC, Kafka) -> attach business logic to this transport

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	httpClient := client.NewHTTPClient(aggregatorEndpoint)
	//grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	//if err != nil {
	//	log.Fatal(err)
	//}

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, httpClient)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
