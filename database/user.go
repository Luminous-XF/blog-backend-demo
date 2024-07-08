package database

import (
	"blog-backend/util"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Id       int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username string `gorm:"column:username;unique"`
	Nickname string `gorm:"column:nickname"`
	Password string `gorm:"column:password"`
	Salt     string `gorm:"column:salt"`
	Email    string `gorm:"column:email;unique"`
}

// TableName 显示指定表名
func (User) TableName() string {
	return "user"
}

func GetUserByUsername(username string) (*User, error) {
	db := GetBlogDBConnection()

	var user User
	if err := db.Select([]string{"id", "username", "password"}).Where("username = ?", username).First(&user).Error; err != nil {
		// 如果查询用户不存在不输出错误日志, 只有系统错误才输出日志
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			util.Logger.Errorf("get password of user %s failed: %s", username, err)
		}
		return nil, err
	}

	return &user, nil
}
