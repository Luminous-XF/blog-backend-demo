package service

import (
	"blog-backend/common/error_code"
	"blog-backend/database"
	"blog-backend/model"
	"blog-backend/model/request"
	"blog-backend/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func LoginByUsernameAndPassword(loginForm request.LoginByUsernameAndPasswordRequest) (code error_code.ErrorCode) {
	if code = checkUsername(loginForm.Username); !error_code.IsSuccess(code) {
		return code
	}

	if code = checkPassword(loginForm.Username, loginForm.Password); !error_code.IsSuccess(code) {
		return code
	}

	return error_code.SUCCESS
}

func checkUsername(username string) error_code.ErrorCode {
	length := len(username)

	// 判断用户名是否为空
	if length == 0 {
		return error_code.UsernameCanNotBlank
	}

	// 检查用户名长度
	if length < 3 || length > 16 {
		return error_code.IllegalUsernameLength
	}

	// 判断用户名是否存在
	if ok, err := isUsernameExist(username); !ok && err != nil {
		return error_code.UsernameIsNotExist
	}

	return error_code.SUCCESS
}

func checkPassword(username, password string) error_code.ErrorCode {
	length := len(password)

	// 判断密码是否为空
	if length == 0 {
		return error_code.PasswordCanNotBlank
	}

	// 检查密码长度
	if length < 8 || length > 16 {
		return error_code.IllegalPasswordLength
	}

	// 校验密码
	if !verifyPassword(username, password) {
		return error_code.PasswordVerifyFail
	}

	return error_code.SUCCESS
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
