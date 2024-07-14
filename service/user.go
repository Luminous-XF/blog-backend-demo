package service

import (
	"blog-backend/database"
	"blog-backend/model"
	"blog-backend/model/request"
	"blog-backend/model/response"
	"blog-backend/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func LoginByUsernameAndPassword(loginForm request.LoginByUsernamePassword) response.ErrorCode {
	if code := checkUsername(loginForm.Username); !response.IsSuccess(code) {
		return code
	}

	if code := checkPassword(loginForm.Username, loginForm.Password); !response.IsSuccess(code) {
		return code
	}

	return response.LoginSuccess
}

func checkUsername(username string) response.ErrorCode {
	length := len(username)

	// 判断用户名是否为空
	if length == 0 {
		return response.UsernameCanNotBlank
	}

	// 检查用户名长度
	if length < 3 || length > 16 {
		return response.IllegalUsernameLength
	}

	// 判断用户名是否存在
	if ok, err := isUsernameExist(username); !ok && err != nil {
		return response.UsernameIsNotExist
	}

	return response.SUCCESS
}

func checkPassword(username, password string) response.ErrorCode {
	length := len(password)

	// 判断密码是否为空
	if length == 0 {
		return response.PasswordCanNotBlank
	}

	// 检查密码长度
	if length < 8 || length > 16 {
		return response.IllegalPasswordLength
	}

	// 校验密码
	if !verifyPassword(username, password) {
		return response.PasswordVerifyFail
	}

	return response.SUCCESS
}

func isUsernameExist(username string) (bool, *model.User) {
	user, err := database.GetUserByUsername(username)
	return (err == nil || !errors.Is(err, gorm.ErrRecordNotFound)) && user != nil, user
}

// 校验密码
// 先对用户输入的密码拼接盐值, 再将结果的 MD5 值与数据库中的 Password 进行比较
func verifyPassword(username, password string) bool {
	user, _ := database.GetUserByUsername(username)
	passwordMd5 := utils.Md5(password + user.Salt)
	return passwordMd5 == user.Password
}
