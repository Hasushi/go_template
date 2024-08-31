package input_port

import "go_template/domain/entity"

type UserCreateArgs struct {
	Email string
	Password string
	UserType string
	Name string
}

type UserUpdateArgs struct {
	UserID string
	UserType string
	Name string
}

type UserUpdatePasswordArgs struct {
	UserID string
	NewPassword string
}

type CheckAuthenticationCodeForResetPasswordArgs struct {
	Email              string
	AuthenticationCode string
}

type IUserUseCase interface {
	Authenticate(token string) (string, error)
	AuthenticateForUpdateEmail(token string) (string, error)
	AuthenticateForUpdatePassword(token string) (string, error)
	Login(email, password string) (entity.User, string, error)
	Create(args UserCreateArgs) (entity.User, error)
	CreateUserWithDetail(user entity.User) error
	Delete(userID string) (entity.User, error)
	FindByID(userID string) (entity.User, error)
	Search(query, userType string, skip int, limit int) ([]*entity.User, int, error)
	Update(args UserUpdateArgs) (entity.User, error)
	UpdateEmailRequest(userID, newEmail string) (entity.User, error)
	UpdateEmail(email string) (entity.User, error)
	UpdatePassword(args UserUpdatePasswordArgs) error
	SendForgetPasswordMail(emailAddress string) (string, error)
	CheckAuthenticationCodeForResetPassword(args CheckAuthenticationCodeForResetPasswordArgs) (entity.User, string, error)
}