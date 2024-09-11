package output_port

import (
	"go_template/domain/entity"
	"time"
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