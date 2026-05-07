package utils

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// IsValidEmail returns true if the value is a valid email address.
func IsValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	return email != "" && emailRegex.MatchString(email)
}

// IsNonEmpty returns true if the string contains non-whitespace characters.
func IsNonEmpty(value string) bool {
	return strings.TrimSpace(value) != ""
}

// IsPositivePrice returns true if price is greater than zero.
func IsPositivePrice(price float64) bool {
	return price > 0
}
