package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

type Producer struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaProducer(brokerAddress, topic string) *Producer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}
	return &Producer{
		writer: writer,
		topic:  topic,
	}
}

func (k *Producer) SendWelcomeEmail(email, nickname string) error {
	msg := map[string]string{
		"email":    email,
		"nickname": nickname,
	}
	value, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	return k.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(email),
		Value: value,
		Time:  time.Now(),
	})
}
