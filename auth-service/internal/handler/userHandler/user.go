package userHandler

import (
	"auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	userService service.UserServiceInterface
	producer    service.EventProducerInterface
}

func NewUserHandler(userService service.UserServiceInterface, producer service.EventProducerInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
		producer:    producer}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn("invalid register request", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.userService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = h.producer.SendWelcomeEmail(req.Email, req.Nickname)
	if err != nil {
		slog.Error("error send welcome email", "err", err)
		return
	}
	slog.Info("user registration successful", "email", req.Email)
	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func (h *UserHandler) Update(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn("invalid update request", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.userService.Update(req.Email, req.Nickname, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	slog.Info("user update successful", "email", req.Email)
	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}
