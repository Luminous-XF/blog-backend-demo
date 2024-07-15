package model

import "github.com/google/uuid"

type User struct {
	Model
	UUID     uuid.UUID `gorm:"column:uuid"`
	Username string    `gorm:"column:username;unique"`
	Nickname string    `gorm:"column:nickname"`
	Password string    `gorm:"column:password"`
	Salt     string    `gorm:"column:salt"`
	Email    string    `gorm:"column:email;unique"`
}

func (User) TableName() string {
	return "user"
}
