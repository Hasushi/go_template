package entity_const

var (
	ErrInvalidUserType = NewValidationErrorFromMsg("invalid user type")
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

