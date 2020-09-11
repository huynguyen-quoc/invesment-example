package main

import (
	"fmt"
	"github.com/huynguyen-quoc/streams/kafka"
)

const (
	Broker1 = "localhost:9092"
	Broker2 = "localhost:9093"
	Broker3 = "localhost:9094"
)

var (
	KafkaCluster = []string{Broker1, Broker2, Broker3}
)

func main() {
	_, err := kafka.NewConsumer(&kafka.ConsumerConfig{
		Brokers:         KafkaCluster,
		ConsumerGroupID: "test",
		Topic:           "test",
		ClientID:        "test",
		InitOffset:      kafka.OffsetOldest,
	})

	if err != nil {
		fmt.Printf("error [%v]", err)
	}

	fmt.Printf("success")

}
