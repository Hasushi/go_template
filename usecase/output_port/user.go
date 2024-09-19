package output_port

import (
	"go_template/domain/entity"
	"go_template/domain/entity_const"
	"time"
)

var (
	TokenScopeGeneral                 = "general"
	TokenGeneralExpireDuration        = 7 * 24 * time.Hour // 7 days
	TokenScopeUpdateEmail             = "updateEmail"
	TokenEmailUpdateExpireDuration    = 24 * time.Hour // 1 day
	TokenScopeUpdatePassword          = "updatePassword"
	TokenChangePasswordExpireDuration = 24 * time.Hour // 1 day
	TokenForgetPasswordExpireDuration = 1 * time.Hour  // 1 hour
	ErrUnknownScope                   = entity_const.NewValidationErrorFromMsg("unknown scope")
	ErrTokenExpired                   = entity_const.NewValidationErrorFromMsg("token expired")
	ErrTokenIssuedFutureTime          = entity_const.NewValidationErrorFromMsg("token issued future time")
	ErrTokenScopeInvalid              = entity_const.NewValidationErrorFromMsg("token scope invalid")
	ErrUserNotFound                   = entity_const.NewNotFoundErrorFromMsg("user not found")
)

type UserAuth interface {
	Authenticate(token string) (string, error)
	AuthenticateForUpdateEmail(token string) (string, error)
	AuthenticateForUpdatePassword(token string) (string, error)
	CheckPassword(user entity.User, password string) error
	CheckAuthenticationCodeForResetPassword(hashedAc, ac string) error
	HashPassword(password string) (string, error)
	IssueUserToken(user entity.User, issueAt time.Time) (string, error)
	IssueUserTokenForUpdateEmail(user entity.User, issuedAt time.Time) (string, error)
	IssueUserTokenForUpdatePassword(user entity.User, issuedAt time.Time) (string, error)
	GenerateInitialPassword(length int) (string, error)
}