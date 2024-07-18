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
	"bytes"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"html/template"
	"net/smtp"
	"time"
)

func LoginByUsernameAndPassword(requestData request.LoginByUsernameAndPasswordRequest) (responseData *response.LoginResponse, code error_code.ErrorCode) {
	user, isExist := IsUsernameExist(requestData.Username)
	if !isExist {
		return nil, error_code.UsernameIsNotExist
	}

	// 校验密码
	passwordMd5 := utils.Md5(requestData.Password + user.Salt)
	if passwordMd5 != user.Password {
		return nil, error_code.PasswordVerifyFailed
	}

	tokenStr, code := CreateToken(user)
	if !error_code.IsSuccess(code) {
		return nil, code
	}

	responseData = &response.LoginResponse{
		User: response.UserResponse{
			UUID:           user.UUID,
			Username:       user.Username,
			Nickname:       user.Nickname,
			Email:          user.Email,
			AvatarImageURL: user.AvatarImageURL,
		},
		Token: tokenStr,
	}

	return responseData, error_code.SUCCESS
}

func SendVerifyCodeWithEmail(requestData request.SendVerifyCodeWithEmailRequest, requestId string) (responseData *response.SendVerifyCodeWithEmailResponse, code error_code.ErrorCode) {
	if _, ok := IsUsernameExist(requestData.Username); ok {
		return nil, error_code.UsernameAlreadyExists
	}

	if _, ok := IsEmailExist(requestData.Email); ok {
		return nil, error_code.EmailAlreadyInUse
	}

	// 生成验证码
	verifyCode := utils.MakeStr(6, utils.DigitAlpha)

	// 将验证码存入 Redis

	// 生成邮件
	t, err := template.ParseFiles("templates/verify-code.html")
	if err != nil {
		global.Logger.Errorf("Create email template failed! err: %s", err)
	}

	var emailBody bytes.Buffer
	err = t.Execute(&emailBody, struct {
		Username   string
		Email      string
		VerifyCode string
	}{
		Username:   requestData.Username,
		Email:      requestData.Email,
		VerifyCode: verifyCode,
	})
	if err != nil {
		global.Logger.Errorf("Create email template failed! err: %s", err)
	}

	e := &email.Email{
		To:      []string{requestData.Email},
		From:    global.CONFIG.EmailConfig.Addr,
		Subject: "XOJ VerifyCode",
		HTML:    []byte(emailBody.String()),
	}
	emailAuth := smtp.PlainAuth(
		"",
		global.CONFIG.EmailConfig.Addr,
		global.CONFIG.EmailConfig.LicenseKey,
		"smtp.qq.com",
	)
	// 发送邮件
	err = e.Send("smtp.qq.com:587", emailAuth)
	if err != nil {
		global.Logger.Errorf("Send email failed! err: %s", err)
	}

	responseData = &response.SendVerifyCodeWithEmailResponse{
		RequestID: requestId,
	}
	return responseData, error_code.SUCCESS
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

func IsEmailExist(email string) (*model.User, bool) {
	user, err := database.GetUserByEmail(email)
	return user, (err == nil || !errors.Is(err, gorm.ErrRecordNotFound)) && user != nil
}
