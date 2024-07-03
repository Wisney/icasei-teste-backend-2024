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

func Message(data []byte) kafka.Message{
	return kafka.Message{
		Value: data,
	}
}