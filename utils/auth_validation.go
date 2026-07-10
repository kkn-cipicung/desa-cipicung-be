package utils

import (
	"errors"
	"fmt"
	"strings"
)

const (
	minAuthNameLength     = 2
	maxAuthNameLength     = 100
	minAuthUsernameLength = 3
	maxAuthUsernameLength = 32
	minAuthPasswordLength = 8
	maxAuthPasswordLength = 72
)

var ErrInvalidAuthPayload = errors.New("invalid auth payload")

func NormalizeAuthName(name string) string {
	return strings.Join(strings.Fields(name), " ")
}

func NormalizeAuthUsername(username string) string {
	return strings.ToLower(strings.TrimSpace(username))
}

func ValidateAuthRegisterPayload(name, username, password string) error {
	if err := ValidateAuthName(name); err != nil {
		return err
	}

	if err := ValidateAuthUsername(username); err != nil {
		return err
	}

	return ValidateAuthPassword(password)
}

func ValidateAuthLoginPayload(username, password string) error {
	if err := ValidateAuthUsername(username); err != nil {
		return err
	}

	if password == "" {
		return fmt.Errorf("%w: password is required", ErrInvalidAuthPayload)
	}

	if len(password) > maxAuthPasswordLength {
		return fmt.Errorf("%w: password must be at most %d characters", ErrInvalidAuthPayload, maxAuthPasswordLength)
	}

	return nil
}

func ValidateAuthName(name string) error {
	if len(name) < minAuthNameLength {
		return fmt.Errorf("%w: name must be at least %d characters", ErrInvalidAuthPayload, minAuthNameLength)
	}

	if len(name) > maxAuthNameLength {
		return fmt.Errorf("%w: name must be at most %d characters", ErrInvalidAuthPayload, maxAuthNameLength)
	}

	return nil
}

func ValidateAuthUsername(username string) error {
	if len(username) < minAuthUsernameLength {
		return fmt.Errorf("%w: username must be at least %d characters", ErrInvalidAuthPayload, minAuthUsernameLength)
	}

	if len(username) > maxAuthUsernameLength {
		return fmt.Errorf("%w: username must be at most %d characters", ErrInvalidAuthPayload, maxAuthUsernameLength)
	}

	for _, char := range username {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '_' || char == '.' {
			continue
		}
		return fmt.Errorf("%w: username can only contain lowercase letters, numbers, underscores, and dots", ErrInvalidAuthPayload)
	}

	return nil
}

func ValidateAuthPassword(password string) error {
	// if strings.TrimSpace(password) != password {
	// 	return fmt.Errorf("%w: password must not start or end with whitespace", ErrInvalidAuthPayload)
	// }

	// if len(password) < minAuthPasswordLength {
	// 	return fmt.Errorf("%w: password must be at least %d characters", ErrInvalidAuthPayload, minAuthPasswordLength)
	// }

	// if len(password) > maxAuthPasswordLength {
	// 	return fmt.Errorf("%w: password must be at most %d characters", ErrInvalidAuthPayload, maxAuthPasswordLength)
	// }

	// var hasLetter, hasNumber, hasSymbol bool
	// for _, char := range password {
	// 	switch {
	// 	case unicode.IsLetter(char):
	// 		hasLetter = true
	// 	case unicode.IsNumber(char):
	// 		hasNumber = true
	// 	case unicode.IsPunct(char) || unicode.IsSymbol(char):
	// 		hasSymbol = true
	// 	}
	// }

	// if !hasLetter || !hasNumber || !hasSymbol {
	// 	return fmt.Errorf("%w: password must contain letters, numbers, and symbols", ErrInvalidAuthPayload)
	// }

	return nil
}
