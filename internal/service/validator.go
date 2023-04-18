package service

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
)

func ValidateSting(field, value string, minLen, maxLen int) error {
	if n := len(value); n < minLen || n > maxLen {
		return fmt.Errorf("%v must contain from %v-%v characters", field, minLen, maxLen)
	}
	return nil
}

func ValidateUsername(value string) (err error) {
	if err = ValidateSting("username", value, 3, 12); err != nil {
		return
	}
	if !isValidUsername(value) {
		return fmt.Errorf("username must contain only lowercase letters, digits, or underscore")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateSting("password", value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateSting("email", value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email")
	}
	return nil

}
