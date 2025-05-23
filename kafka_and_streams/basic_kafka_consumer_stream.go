package main

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	// Create Kafka client
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics("input-topic"),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Stream processing loop
	ctx := context.Background()
	for {
		fetches := client.PollFetches(ctx)
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Printf("Received message: %s\n", string(record.Value))
		}
	}
}

