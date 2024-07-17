package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// 密码仅包含:英文字母、数字、下滑线
var usernameCharset validator.Func = func(fieldLevel validator.FieldLevel) bool {
	username := fieldLevel.Field().String()
	regex := `^[a-zA-Z0-9_-]+$`
	matched, _ := regexp.MatchString(regex, username)
	return matched
}

// 密码仅包含:英文字母、数字、下滑线
var passwordCharset validator.Func = func(fieldLevel validator.FieldLevel) bool {
	password := fieldLevel.Field().String()
	regex := `^[a-zA-Z\d!@#$%^&*(),.?":{}|<>]+$`
	matched, _ := regexp.MatchString(regex, password)
	return matched
}

func initValidator() (validate *validator.Validate) {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		if err := validate.RegisterValidation("username-charset", usernameCharset); err != nil {
			return nil
		}
		if err := validate.RegisterValidation("password-charset", passwordCharset); err != nil {
			return nil
		}
	}

	return validate
}
