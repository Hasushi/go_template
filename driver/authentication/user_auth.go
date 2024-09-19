package authentication

import (
	"crypto/rand"
	"go_template/domain/entity"
	"go_template/usecase/output_port"
	"math/big"
	"time"
)

type UserAuth struct{}

func NewUserAuth() output_port.UserAuth {
	return &UserAuth{}
}

func (a *UserAuth) CheckPassword(u entity.User, password string) error {
	return CheckBcryptPassword(u.HashedPassword, password)
}

func (a *UserAuth) CheckAuthenticationCodeForResetPassword(hashedAc, ac string) error {
	return CheckBcryptPassword(hashedAc, ac)
}

func (a *UserAuth) IssueUserToken(user entity.User, issuedAt time.Time) (string, error) {
	return IssueUserToken(user.UserID, issuedAt, []string{output_port.TokenScopeGeneral})
}

func (a *UserAuth) IssueUserTokenForUpdateEmail(user entity.User, issuedAt time.Time) (string, error) {
	return IssueUserToken(user.UserID, issuedAt, []string{output_port.TokenScopeUpdateEmail})
}

func (a *UserAuth) IssueUserTokenForUpdatePassword(user entity.User, issuedAt time.Time) (string, error) {
	return IssueUserToken(user.UserID, issuedAt, []string{output_port.TokenScopeUpdatePassword})
}

func (a *UserAuth) HashPassword(password string) (string, error) {
	hp, err := HashBcryptPassword(password)
	if err != nil {
		return "", err
	}
	return hp, nil
}

func (a *UserAuth) Authenticate(token string) (string, error) {
	return VerifyUserToken(token, []string{output_port.TokenScopeGeneral})
}

func (a *UserAuth) AuthenticateForUpdateEmail(token string) (string, error) {
	return VerifyUserToken(token, []string{output_port.TokenScopeUpdateEmail})
}

func (a *UserAuth) AuthenticateForUpdatePassword(token string) (string, error) {
	return VerifyUserToken(token, []string{output_port.TokenScopeUpdatePassword})
}

func (a *UserAuth) GenerateInitialPassword(length int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}
