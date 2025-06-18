package main

import (
	"auth-service/internal/handler"
	"auth-service/internal/handler/authHandler"
	"auth-service/internal/handler/userHandler"
	"auth-service/internal/service/auth"
	"auth-service/internal/service/user"
	"auth-service/internal/storage/postgres"
	"auth-service/pkg/tokenManager/jwtManager"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	db, err := postgres.InitPostgres()
	if err != nil {
		log.Fatal("failed to connect DB:", err)
	}
	repo := postgres.NewPostgresRepository(db)

	tokenManager := jwtManager.NewJwtTokenManager(os.Getenv("JWT_SECRET"))

	authService := auth.NewAuthService(repo, tokenManager)
	userService := user.NewUserService(repo)

	aHandler := authHandler.NewAuthHandler(authService)
	uHandler := userHandler.NewUserHandler(userService)

	handler.RegisterRoutes(r, aHandler, uHandler)

	// Запуск
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
