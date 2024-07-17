package service

import (
	"blog-backend/common/error_code"
	"blog-backend/database"
	"blog-backend/global"
	"blog-backend/model"
	"blog-backend/model/request"
	"blog-backend/model/response"
	"blog-backend/model/token"
	"blog-backend/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func LoginByUsernameAndPassword(formData request.LoginByUsernameAndPasswordRequest) (loginResponse *response.LoginResponse, code error_code.ErrorCode) {
	user, isExist := IsUsernameExist(formData.Username)
	if !isExist {
		return nil, error_code.UsernameIsNotExist
	}

	// 校验密码
	passwordMd5 := utils.Md5(formData.Password + user.Salt)
	if passwordMd5 != user.Password {
		return nil, error_code.PasswordVerifyFailed
	}

	tokenStr, code := CreateToken(user)
	if !error_code.IsSuccess(code) {
		return nil, code
	}

	loginResponse = &response.LoginResponse{
		User: response.UserResponse{
			UUID:           user.UUID,
			Username:       user.Username,
			Nickname:       user.Nickname,
			Email:          user.Email,
			AvatarImageURL: user.AvatarImageURL,
		},
		Token: tokenStr,
	}

	return loginResponse, error_code.SUCCESS
}

func CreateToken(user *model.User) (tokenStr string, code error_code.ErrorCode) {
	j := &JWT{
		SigningKey: []byte(global.CONFIG.JWTConfig.SigningKey),
	}

	claims := token.CustomClaims{
		UUID:     user.UUID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(global.CONFIG.JWTConfig.ExpiresTime))),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Luminous",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenStr, err := j.GenToken(claims)
	if err != nil {
		return "", error_code.AuthTokenCreateFailed
	}

	return tokenStr, error_code.SUCCESS
}

func IsUsernameExist(username string) (*model.User, bool) {
	user, err := database.GetUserByUsername(username)
	return user, (err == nil || !errors.Is(err, gorm.ErrRecordNotFound)) && user != nil
}
