package validations

import (
	"regexp"

	"go_template/domain/entity_const"	
)

var (
	ErrInvalidUserID = entity_const.NewValidationErrorFromMsg("invalid user id")
	ErrInvalidUserDetailID = entity_const.NewValidationErrorFromMsg("invalid user detail id")
	ErrInvalidName = entity_const.NewValidationErrorFromMsg("invalid name")
)

var (
	ErrTooLongEmail = entity_const.NewValidationErrorFromMsg("email address is too long")
	ErrInvalidEmail = entity_const.NewValidationErrorFromMsg("email address is invalid")
)

func ValidateEmail(email string) error {
	if len(email) > 256 {
		return ErrTooLongEmail
	}

	const emailPattern string = `^[a-zA-Z0-9_+-.]+@[a-z0-9-.]+\.[a-z]+$`
	var emailRegexp *regexp.Regexp = regexp.MustCompile(emailPattern)
	if !emailRegexp.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

var (
	ErrTooShortPassword = entity_const.NewValidationErrorFromMsg("password is too short")
	ErrTooLongPassword  = entity_const.NewValidationErrorFromMsg("password is too long")
	ErrInvalidPassword  = entity_const.NewValidationErrorFromMsg("password is invalid")
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrTooShortPassword
	}

	if len(password) > 256 {
		return ErrTooLongPassword
	}

	const pattern string = `^[a-zA-Z0-9_+-.]+$`
	var passwordRegexp *regexp.Regexp = regexp.MustCompile(pattern)
	if !passwordRegexp.MatchString(password) {
		return ErrInvalidPassword
	}

	return nil
}
