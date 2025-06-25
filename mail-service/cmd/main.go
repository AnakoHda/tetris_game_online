package main

import (
	"context"
	"log/slog"
	"mail_service/internal/config"
	"mail_service/internal/handler/leaderUpdateConsumer"
	"mail_service/internal/handler/welcomeEmailConsumer"
	"mail_service/internal/service/mailSender"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := config.ParseEnvironment(); err != nil {
		slog.Error("parse environment", "error", err)
		os.Exit(1)
	}
	sender := mailSender.NewMailSender(
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_HOST"),
		os.Getenv("MAIL_PORT"))

	kafkaBrokers := []string{os.Getenv("KAFKA_BROKERS")}
	kafkaWelcomeTopic := os.Getenv("KAFKA_HALLO_TOPIC")
	kafkaScoreUpdateTopic := os.Getenv("KAFKA_SCORE_UPDATE_TOPIC")
	kafkaGroupID := os.Getenv("KAFKA_GROUP_ID")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	kafkaWelcomeMessage := welcomeEmailConsumer.NewWelcomeEmailConsumer(sender, kafkaBrokers, kafkaWelcomeTopic, kafkaGroupID)
	go kafkaWelcomeMessage.Start(ctx)
	slog.Info("welcome consumer started")
	kafkaLeaderUpdate := leaderUpdateConsumer.NewLeaderUpdateConsumer(sender, kafkaBrokers, kafkaScoreUpdateTopic, kafkaGroupID)
	go kafkaLeaderUpdate.Start(ctx)
	slog.Info("leader update consumer started")

	slog.Info("mail-service Started")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	slog.Info("mail-service Stop")
}
