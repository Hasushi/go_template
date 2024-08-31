package schema

import "go_template/domain/entity"

const TokenType = "Bearer"

type LoginResUser struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
}

type LoginRes struct {
	AccessToken string       `json:"accessToken"`
	TokenType   string       `json:"tokenType"`
	User        LoginResUser `json:"user"`
}

type SendForgetPasswordMailRes struct {
	Email string `json:"email"`
}

type CheckAuthenticationCodeForResetPasswordRes struct {
	AccessToken string        `json:"accessToken"`
	TokenType   string        `json:"tokenType"`
	User        *LoginResUser `json:"user"`
}

func LoginResUserFromEntity(user *entity.User) *LoginResUser {
	if user == nil {
		return nil
	}

	return &LoginResUser{
		UserId: user.UserID,
		Email:  user.Email,
	}
}

func CheckAuthenticationCodeForResetPasswordResFromEntity(user *entity.User, token string) *CheckAuthenticationCodeForResetPasswordRes {
	if user == nil {
		return nil
	}

	return &CheckAuthenticationCodeForResetPasswordRes{
		AccessToken: token,
		TokenType:   TokenType,
		User:        LoginResUserFromEntity(user),
	}
}
