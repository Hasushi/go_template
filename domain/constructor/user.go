package constructor

import (
	"go_template/domain/entity"
	"go_template/domain/entity_const"
)


type NewUserDetailCreateArgs struct {
	UserDetailID string
	UserID string
	Name string
}

func NewUserDetailCreate(args NewUserDetailCreateArgs) (entity.UserDetail, error) {
	if args.UserDetailID == "" {
		return entity.UserDetail{}, entity_const.ErrInvalidUserDetailID
	}
	if args.UserID == "" {
		return entity.UserDetail{}, entity_const.ErrInvalidUserID
	}
	if args.Name == "" {
		return entity.UserDetail{}, entity_const.ErrInvalidName
	}
	
	return entity.UserDetail{
		UserDetailID: args.UserDetailID,
		UserID: args.UserID,
		Name: args.Name,
		}, nil
}
	
type NewUserCreateArgs struct {
	UserID string
	Email string
	HashedPassword string
	UserType string
}

func NewUserCreate(args NewUserCreateArgs) (entity.User, error) {
	if args.UserID == "" {
		return entity.User{}, entity_const.ErrInvalidUserID
	}
	if args.Email == "" {
		return entity.User{}, entity_const.ErrInvalidEmail
	}
	if args.HashedPassword == "" {
		return entity.User{}, entity_const.ErrInvalidHashedPassword
	}
	if _, err := entity_const.UserTypeFromString(args.UserType); err != nil {
		return entity.User{}, err
	}
	
	return entity.User{
		UserID: args.UserID,
		Email: args.Email,
		HashedPassword: args.HashedPassword,
		UserType: entity_const.UserType(args.UserType),
		}, nil
}