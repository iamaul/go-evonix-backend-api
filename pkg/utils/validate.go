package utils

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// Use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct Validate struct fields
func ValidateStruct(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}

// IsValidPhoneNumber validate phone number
func IsValidPhoneNumber(phoneNumber string) bool {
	regexPhoneNumber := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return regexPhoneNumber.MatchString(phoneNumber)
}

// IsValidGender validate gender
func IsValidGender(gender string) bool {
	genderLower := strings.ToLower(gender)
	if genderLower == "male" || genderLower == "female" {
		return true
	}
	return false
}

// IsValidBirthDate validate birthdate
func IsValidBirthDate(birthdate string) bool {
	const layoutISO = "2006-01-02"
	t, err := time.Parse(layoutISO, birthdate)
	if err != nil {
		return false
	}
	if t.IsZero() {
		return false
	}
	return true
}
