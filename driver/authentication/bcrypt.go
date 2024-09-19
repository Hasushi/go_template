package authentication

import (
	"encoding/base64"
	"go_template/domain/entity_const"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrBase64DecodeFailed = entity_const.NewValidationErrorFromMsg("failed to decode in base64")
	ErrPasswordHashFailed = entity_const.NewValidationErrorFromMsg("failed to hash password")
	ErrUnexpected         = entity_const.NewValidationErrorFromMsg("unexpected error occurred in password comparing")
	ErrWrongPassword      = entity_const.NewValidationErrorFromMsg("password authentication failed")
	ErrAuthenticationCode = entity_const.NewValidationErrorFromMsg("authenticationCode authentication failed")
)

func bcryptHash(txt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(txt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hash), nil
}

func HashBcryptPassword(password string) (string, error) {
	hash, err := bcryptHash(password)
	if err != nil {
		return "", ErrPasswordHashFailed
	}

	return hash, nil
}

func CheckBcryptPassword(hashedPassword string, password string) error {
	hb, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return ErrBase64DecodeFailed
	}

	err = bcrypt.CompareHashAndPassword(hb, []byte(password))
	if err == nil {
		return nil
	}
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrWrongPassword
	}
	return ErrUnexpected
}
