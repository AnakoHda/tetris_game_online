package validator

import (
	"errors"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	var (
		upper   = regexp.MustCompile(`[A-Z]`)
		lower   = regexp.MustCompile(`[a-z]`)
		number  = regexp.MustCompile(`[0-9]`)
		special = regexp.MustCompile(`[^a-zA-Z0-9]`)
	)

	switch {
	case !upper.MatchString(password):
		return errors.New("password must contain at least one uppercase letter")
	case !lower.MatchString(password):
		return errors.New("password must contain at least one lowercase letter")
	case !number.MatchString(password):
		return errors.New("password must contain at least one digit")
	case !special.MatchString(password):
		return errors.New("password must contain at least one special character")
	default:
		return nil
	}
}
