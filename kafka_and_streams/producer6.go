package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    broker := "localhost:9092"
    topic := "logs"

    admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": broker})
    if err != nil {
        panic(err)
    }
    defer admin.Close()

    // Получаем метаданные
    metadata, err := admin.GetMetadata(&topic, false, 5000)
    if err != nil {
        panic(err)
    }

    // Выводим информацию о партициях топика
    if topicMeta, ok := metadata.Topics[topic]; ok {
        fmt.Printf("Топик: %s\n", topic)
        for _, p := range topicMeta.Partitions {
            fmt.Printf("Партиция: %d, Лидер: %d, Реплики: %v\n", p.ID, p.Leader, p.Replicas)
        }
    } else {
        fmt.Println("Топик не найден.")
    }
}
