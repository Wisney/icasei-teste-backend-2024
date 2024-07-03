package kafka

import (
	"github.com/segmentio/kafka-go"
)

func Writer() *kafka.Writer{
	writer := kafka.Writer{
		Addr:     kafka.TCP("kafka:29092"),
		Topic:   "go-to-rails",
		Balancer: &kafka.LeastBytes{},
	}
	
	return &writer
}

func Reader() *kafka.Reader{
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:29092"},
		GroupID:   "consumer-group-rails-to-go",
		Topic:     "rails-to-go",
		MaxBytes:  10e6, // 10MB
	})
	
	return reader
}

func Message(data []byte) kafka.Message{
	return kafka.Message{
		Value: data,
	}
}