package entity

import "go_template/domain/entity_const"

type User struct {
	UserID string
	Email string
	EmailToUpdate string
	UserType entity_const.UserType
	HashedPassword string
	UserDetail *UserDetail
}

type UserDetail struct {
	UserDetailID string
	UserID string
	Name string
}