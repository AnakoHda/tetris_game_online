package mailSender

import (
	"fmt"
	"log/slog"
	"mail_service/internal/models"
	"net/smtp"
)

type MailSender struct {
	from     string
	password string
	host     string
	port     string
}

func NewMailSender(from, password, host, port string) *MailSender {
	return &MailSender{
		from:     from,
		password: password,
		host:     host,
		port:     port,
	}
}
func (m *MailSender) SendWelcomeEmail(data models.WelcomeEmail) error {
	subject := "Добро пожаловать!"
	body := fmt.Sprintf("Привет, %s!\n\nСпасибо за регистрацию в нашей игре brick_game_online!", data.Nickname)
	err := m.sendEmail(data.Email, subject, body)
	if err != nil {
		slog.Error("fail to send welcome email", "to", data.Email, "err", err)
		return err
	}
	slog.Info("sending welcome email successful", "to", data.Email)
	return nil
}
func (m *MailSender) SendLeaderUpdate(data models.LeaderUpdateEmail) error {
	subject := "Лидерборд: Вы больше не на первом месте"
	body := fmt.Sprintf("Привет, %s!\n\nК сожалению, вас сестили с первой позиции."+
		"\nТекущий лидер: %s с %d очками.\n\nПопробуйте снова и верните лидерство!\n",
		data.Nickname, data.NewLeader, data.NewScore)
	err := m.sendEmail(data.Email, subject, body)
	if err != nil {
		slog.Error("fail to send leader update email", "to", data.Email, "err", err)
		return err
	}
	slog.Info("sending leader update email successful", "to", data.Email)
	return nil
}

func (m *MailSender) sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", m.from, m.password, m.host)
	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))
	addr := fmt.Sprintf("%s:%s", m.host, m.port)
	err := smtp.SendMail(addr, auth, m.from, []string{to}, msg)
	if err != nil {
		slog.Error("SMTP error send email", "to", to, "err", err)
	}
	return err
}
