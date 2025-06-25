package welcomeEmailConsumer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log/slog"
	"mail_service/internal/models"
	"mail_service/internal/service"
)

type WelcomeEmailConsumer struct {
	reader      *kafka.Reader
	mailService service.MailService
	brokers     []string
	topic       string
}

func NewWelcomeEmailConsumer(mailService service.MailService, brokers []string, topic, groupID string) *WelcomeEmailConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})
	return &WelcomeEmailConsumer{
		reader:      reader,
		mailService: mailService,
		brokers:     brokers,
		topic:       topic,
	}
}
func (c *WelcomeEmailConsumer) Start(ctx context.Context) {
	defer c.reader.Close()
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			slog.Error("failed to read message", "err", err)
			continue
		}
		var data models.WelcomeEmail
		if err := json.Unmarshal(m.Value, &data); err != nil {
			slog.Error("invalid JSON payload", "err", err)
			continue
		}
		_ = c.mailService.SendWelcomeEmail(data)
	}
}
