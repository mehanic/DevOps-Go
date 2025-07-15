package main

import (
    "context"
    "log"
    "strings"

    "github.com/segmentio/kafka-go"
)

func main() {
    ctx := context.Background()

    // Конфиг для чтения из топика "input-topic"
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "input-topic",
        GroupID: "example-group",
    })

    // Конфиг для записи в топик "output-topic"
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "output-topic",
    })

    defer reader.Close()
    defer writer.Close()

    for {
        msg, err := reader.ReadMessage(ctx)
        if err != nil {
            log.Fatal("Error reading message: ", err)
        }

        value := string(msg.Value)
        // Фильтрация сообщений, как в Kafka Streams (например, только содержащие "important")
        if strings.Contains(value, "important") {
            err = writer.WriteMessages(ctx, kafka.Message{
                Key:   msg.Key,
                Value: []byte(value),
            })
            if err != nil {
                log.Println("Error writing message: ", err)
            } else {
                log.Printf("Forwarded message: %s\n", value)
            }
        } else {
            log.Printf("Filtered out message: %s\n", value)
        }
    }
}
