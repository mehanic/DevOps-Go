package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    // Создание Producer'а
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    // Канал для асинхронной доставки сообщений
    go func() {
        for e := range p.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    fmt.Printf(" Delivery failed: %v\n", ev.TopicPartition)
                } else {
                    fmt.Printf(" Delivered to %v\n", ev.TopicPartition)
                }
            }
        }
    }()

    topic := "test-topic"

    // Отправка сообщения
    err = p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte("Hello Kafka from Golang!"),
    }, nil)

    if err != nil {
        fmt.Println("❌ Produce error:", err)
    }

    // Ждём завершения всех сообщений
    p.Flush(1000)
}
