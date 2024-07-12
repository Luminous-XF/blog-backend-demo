package database

import (
	"blog-backend/global"
	"blog-backend/model"
)

// GetUserByUsername 通过用户名查找用户
func GetUserByUsername(username string) (user *model.User, err error) {
	err = global.GDB.Where("username = ?", username).First(&user).Error
	return user, err
}

// GetUserByID 通过 ID 查找用户
func GetUserByID(id int64) (user *model.User, err error) {
	err = global.GDB.Where("id = ?", id).First(&user).Error
	return user, err
}

// GetUserByUUID 通过 uuid 查找用户
func GetUserByUUID(uuid string) (user *model.User, err error) {
	err = global.GDB.Where("uuid = ?", uuid).First(&user).Error
	return user, err
}

// CreateUser 新建用户
func CreateUser(user *model.User) error {
	return global.GDB.Create(user).Error
}
