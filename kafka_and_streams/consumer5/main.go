package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	topic := "topic-a"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: 0, // read only error logs
	})

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(msg.Key))
		fmt.Println(string(msg.Value))
	}
}