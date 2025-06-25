package leaderUpdateConsumer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log/slog"
	"mail_service/internal/models"
	"mail_service/internal/service"
)

type LeaderUpdateConsumer struct {
	reader      *kafka.Reader
	mailService service.MailService
	brokers     []string
	topic       string
}

func NewLeaderUpdateConsumer(mailService service.MailService, brokers []string, topic, groupID string) *LeaderUpdateConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})
	return &LeaderUpdateConsumer{
		reader:      reader,
		mailService: mailService,
		brokers:     brokers,
		topic:       topic,
	}
}

func (c *LeaderUpdateConsumer) Start(ctx context.Context) {
	defer c.reader.Close()
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			slog.Error("failed to read message", "err", err)
			continue
		}
		var data models.LeaderUpdateEmail
		if err := json.Unmarshal(m.Value, &data); err != nil {
			slog.Error("invalid JSON payload", "err", err)
			continue
		}
		_ = c.mailService.SendLeaderUpdate(data)
	}
}
