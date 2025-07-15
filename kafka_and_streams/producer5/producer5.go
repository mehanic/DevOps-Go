package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	topic := "topic-a"

	conn, err := kafka.DialContext(ctx, "tcp", "localhost:9092")
	if err != nil {
		panic(err)
	}

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
	})
	if err != nil {
		panic(err)
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &CustomBalancer{},
	})

	err = writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("error"),
			Value: []byte("some error happened in the application"),
		},
		kafka.Message{
			Key:   []byte("warn"),
			Value: []byte("this is a warning"),
		})
	if err != nil {
		panic(err)
	}

	fmt.Println("written messages successfully")
}

type CustomBalancer struct {
	baseBalancer kafka.RoundRobin
}

func (b *CustomBalancer) Balance(msg kafka.Message, partitions ...int) (partition int) {
	// only error logs are sent to first partition
	if string(msg.Key) == "error" {
		return 0
	}

	// everything else gets round robined to other paritions
	return b.baseBalancer.Balance(msg, partitions[1:]...)
}