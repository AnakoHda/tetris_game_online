package handler

import (
	"auth-service/internal/handler/authHandler"
	"auth-service/internal/handler/userHandler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, auth *authHandler.AuthHandler, user *userHandler.UserHandler) {
	r.POST("/auth/login", auth.Login)
	r.GET("/auth/validate", auth.ValidateToken)

	r.POST("/users/register", user.Register)
	r.POST("/users/update", user.Update)
}
