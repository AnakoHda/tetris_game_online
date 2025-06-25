package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/handler/authHandler"
	"auth-service/internal/handler/userHandler"
	"auth-service/internal/service/auth"
	"auth-service/internal/service/kafkaProduser"
	"auth-service/internal/service/user"
	"auth-service/internal/storage/postgres"
	"auth-service/pkg/tokenManager/jwtManager"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func main() {
	r := gin.Default()

	if err := config.ParseEnvironment(); err != nil {
		slog.Error("parse environment", "error", err)
		os.Exit(1)
	}

	db, err := postgres.InitPostgres(os.Getenv("DB_URL"))
	if err != nil {
		slog.Error("init postgres", "error", err)
		os.Exit(1)
	}
	repo := postgres.NewPostgresRepository(db)

	tokenManager := jwtManager.NewJwtTokenManager(os.Getenv("JWT_SECRET"))

	broker := os.Getenv("KAFKA_BROKERS")
	topics := os.Getenv("KAFKA_HALLO_TOPIC")

	producer := kafka.NewKafkaProducer(broker, topics)

	authService := auth.NewAuthService(repo, tokenManager)
	userService := user.NewUserService(repo)

	aHandler := authHandler.NewAuthHandler(authService)
	uHandler := userHandler.NewUserHandler(userService, producer)

	handler.RegisterRoutes(r, aHandler, uHandler)

	if err := r.Run(":8080"); err != nil {
		slog.Error("failed Run ", "error", err)
	}
	slog.Info("auth-service started", "port", 8080)
}
