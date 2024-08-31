package constructor

import (
	"go_template/domain/entity"
	"go_template/domain/entity_const"
	"go_template/domain/validations"
)


type NewUserDetailCreateArgs struct {
	UserDetailID string
	UserID string
	Name string
}

func NewUserDetailCreate(args NewUserDetailCreateArgs) (entity.UserDetail, error) {
	if args.UserDetailID == "" {
		return entity.UserDetail{}, validations.ErrInvalidUserDetailID
	}
	if args.UserID == "" {
		return entity.UserDetail{}, validations.ErrInvalidUserID
	}
	if args.Name == "" {
		return entity.UserDetail{}, validations.ErrInvalidName
	}
	
	return entity.UserDetail{
		UserDetailID: args.UserDetailID,
		UserID: args.UserID,
		Name: args.Name,
		}, nil
}
	
type NewUserCreateArgs struct {
	UserID string
	UserDetailID string
	Email string
	Password string
	HashedPassword string
	UserType string
	Name string
}

func NewUserCreate(args NewUserCreateArgs) (entity.User, error) {
	if err := validations.ValidateEmail(args.Email); err != nil {
		return entity.User{}, err
	}

	if err := validations.ValidatePassword(args.Password); err != nil {
		return entity.User{}, err
	}

	userType, err := entity_const.UserTypeFromString(args.UserType)
	if err != nil {
		return entity.User{}, err
	}

	udCreated, err := NewUserDetailCreate(NewUserDetailCreateArgs{
		UserDetailID: args.UserDetailID,
		UserID: args.UserID,
		Name: args.Name,
	})
	if err != nil {
		return entity.User{}, err
	}
	
	return entity.User{
		UserID: args.UserID,
		Email: args.Email,
		HashedPassword: args.HashedPassword,
		UserType: *userType,
		UserDetail: &udCreated,
	}, nil
}