package utils

import (
	"errors"
	"regexp"
)

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func IsValidPassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}

	var (
		hasUppercase = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLowercase = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber    = regexp.MustCompile(`[0-9]`).MatchString(password)
	)

	if !hasUppercase {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLowercase {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}

	return nil
}
