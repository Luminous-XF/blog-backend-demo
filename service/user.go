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

func LoginByUsernameAndPassword(formData request.LoginByUsernameAndPasswordRequest) (code error_code.ErrorCode) {
	user, isExist := isUsernameExist(formData.Username)
	if !isExist {
		return error_code.UsernameIsNotExist
	}

	// 校验密码
	passwordMd5 := utils.Md5(formData.Password + user.Salt)
	if passwordMd5 != user.Password {
		return error_code.PasswordVerifyFail
	}

	return error_code.SUCCESS
}

func isUsernameExist(username string) (*model.User, bool) {
	user, err := database.GetUserByUsername(username)
	return user, (err == nil || !errors.Is(err, gorm.ErrRecordNotFound)) && user != nil
}
