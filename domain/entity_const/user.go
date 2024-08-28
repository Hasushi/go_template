package entity_const

var (
	ErrInvalidUserType = NewValidationError("invalid user type")
	ErrInvalidUserID = NewValidationError("invalid user id")
	ErrInvalidEmail = NewValidationError("invalid email")
	ErrInvalidHashedPassword = NewValidationError("invalid hashed password")
	ErrInvalidUserDetailID = NewValidationError("invalid user detail id")
	ErrInvalidName = NewValidationError("invalid name")
)

type UserType string

const (
	Admin UserType = "admin"
	User UserType = "user"
)

func (u UserType) String() string {
	return string(u)
}

func UserTypeFromString(s string) (*UserType, error) {
	ret := UserType(s)
	switch ret {
	case Admin, User:
		return &ret, nil
	default:
		return nil, ErrInvalidUserType
	}
}

