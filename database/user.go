package database

import (
	"blog-backend/util"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Uuid     string `gorm:"column:uuid"`
	Id       int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username string `gorm:"column:username;unique"`
	Nickname string `gorm:"column:nickname"`
	Password string `gorm:"column:password"`
	Salt     string `gorm:"column:salt"`
	Email    string `gorm:"column:email;unique"`

	RowVersion    int       `gorm:"column:row_version"`
	RowCreateTime time.Time `gorm:"column:row_create_time;default:null"`
	RowUpdateTime time.Time `gorm:"column:row_update_time;default:null"`
	RowIsDeleted  int       `gorm:"column:row_is_deleted"`
}

// TableName 显示指定表名
func (User) TableName() string {
	return "user"
}

var (
	allUserField = util.GetGormFields(User{})
)

func GetUserByUsername(username string) (*User, error) {
	db := GetBlogDBConnection()

	var user User
	if err := db.Select(allUserField).Where("username = ?", username).First(&user).Error; err != nil {
		// 如果查询用户不存在不输出错误日志, 只有系统错误才输出日志
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			util.Logger.Errorf("get password of user '%s' failed: '%s'", username, err)
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser 创建一个新用户记录
func CreateUser(username, email, password string) {
	db := GetBlogDBConnection()

	uid := uuid.New().String()
	salt := util.GenerateSalt(16)
	password = util.Md5(password + salt)
	user := User{
		Uuid:     uid,
		Username: username,
		Password: password,
		Salt:     salt,
		Email:    email,

		RowCreateTime: time.Now(),
	}

	if err := db.Create(&user).Error; err != nil {
		util.Logger.Errorf("create user '%s' failed: '%s'", username, err)
	} else {
		util.Logger.Infof("create user '%s' successfully id = '%d'", username, user.Id)
	}
}

// DeleteUserByUsername 删除一个用户信息
func DeleteUserByUsername(username string) error {
	db := GetBlogDBConnection()

	err := db.Model(User{}).Where("username = ?", username).Updates(map[string]interface{}{
		"row_is_deleted": 1,
	}).Error

	if err != nil {
		util.Logger.Errorf("delete user '%s' failed: '%s'", username, err)
	} else {
		util.Logger.Infof("delete user '%s' successfully", username)
	}

	return err
}

// DeleteUserByUsername 删除一个用户信息
// func DeleteUserByUsername(username string) {
// 	db := GetBlogDBConnection()
//
// 	if err := db.Where("username = ?", username).Delete(User{}).Error; err != nil {
// 		util.Logger.Errorf("delete user %s failed: %s", username, err)
// 	} else {
// 		util.Logger.Infof("delete user %s successfully", username)
// 	}
// }
