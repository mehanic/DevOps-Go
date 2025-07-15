package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    producer, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    })
    if err != nil {
        panic(err)
    }
    defer producer.Close()

    go func() {
        for e := range producer.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    fmt.Printf("❌ Delivery failed: %v\n", ev.TopicPartition)
                } else {
                    fmt.Printf("✅ Delivered to %v\n", ev.TopicPartition)
                }
            }
        }
    }()

    topic := "logs"

    // Отправка с ключом — Kafka будет использовать хэш ключа для выбора partition
    err = producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Key:            []byte("ERROR"),
        Value:          []byte("Something failed"),
    }, nil)
    if err != nil {
        fmt.Println("Failed to produce (key):", err)
    }

    // Отправка без ключа — используется round-robin между partition'ами
    err = producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte("Info log"),
    }, nil)
    if err != nil {
        fmt.Println("Failed to produce (no key):", err)
    }

    // Ждём, пока все сообщения будут отправлены
    producer.Flush(1000)
}
