package service

import "mail_service/internal/models"

type MailService interface {
	SendWelcomeEmail(data models.WelcomeEmail) error
	SendLeaderUpdate(data models.LeaderUpdateEmail) error
}
