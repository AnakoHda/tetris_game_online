package service

type UserServiceInterface interface {
	Register(req RegisterRequest) error
	Update(email string, newNickname string, newPassword string) error
}
type AuthServiceInterface interface {
	Login(req LoginRequest) (string, error)
	ValidateToken(token string) (bool, error)
}
type EventProducerInterface interface {
	SendWelcomeEmail(email, nickname string) error
}
type RegisterRequest struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
