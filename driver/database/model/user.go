package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID         string `gorm:"primaryKey:not null"`
	Email          string	`gorm:"unique:not null"`
	EmailToUpdate  string	`gorm:"unique"`
	UserType       string	`gorm:"not null"`
	HashedPassword string	`gorm:"not null"`
	UserDetail     *UserDetail	`gorm:"foreignKey:UserID:constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type UserDetail struct {
	UserDetailID string `gorm:"primaryKey:not null"`
	UserID       string	`gorm:"not null"`
	Name         string	`gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}