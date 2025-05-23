// At-Least-Once Producer
producer, _ := kafka.NewProducer(&kafka.ConfigMap{
    "bootstrap.servers": "localhost:9092",
    "acks":              "all",  // Wait for all replicas
    "retries":           3,
})

// At-Least-Once Consumer
consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
    "bootstrap.servers": "localhost:9092",
    "group.id":          "my-group",
    "enable.auto.commit": false, // Commit manually
})

