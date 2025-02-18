// At-Most-Once Producer
producer, _ := kafka.NewProducer(&kafka.ConfigMap{
    "bootstrap.servers": "localhost:9092",
    "acks":              "0", // Fire-and-forget
})

// At-Most-Once Consumer
consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
    "bootstrap.servers": "localhost:9092",
    "group.id":          "my-group",
    "enable.auto.commit": true, // Commit before processing
})

