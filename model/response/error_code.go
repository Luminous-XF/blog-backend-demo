package response

type ErrorCode int // 错误码

//go:generate stringer -type ErrorCode -linecomment

// 20xxx 成功返回 [20000 - 30000)
const (
	SUCCESS      ErrorCode = iota + 20000 // 成功
	LoginSuccess                          // 登录成功
)

// 40xxx 客户端错误
const (
	ERROR          ErrorCode = iota + 40000 // 失败
	ParamBindError                          // 参数信息错误
)

// 41xxx User 模块错误
const (
	UsernameIsNotExist    ErrorCode = iota + 41000 // 用户名不存在
	UsernameCanNotBlank                            // 用户名不能为空
	IllegalUsernameLength                          // 用户名长度应该在3-16字符之间
	PasswordCanNotBlank                            // 密码不能为空
	IllegalPasswordLength                          // 密码长度应该在8-16个字符之间
	PasswordVerifyFail                             // 密码校验失败
)

var errorMsg = map[ErrorCode]string{
	SUCCESS:      "Success!",
	LoginSuccess: "You have successfully logged in!",

	ParamBindError: "There was an error with the parameters provided.",

	UsernameIsNotExist:    "The entered username does not exist.",
	UsernameCanNotBlank:   "The username field cannot be left blank.",
	IllegalUsernameLength: "Username length should be between 8-16 characters.",
	PasswordCanNotBlank:   "The password field cannot be left blank.",
	IllegalPasswordLength: "Password length should be between 8-16 characters.",
	PasswordVerifyFail:    "The password you entered is incorrect. Please try again.",
}

func ErrMsg(code ErrorCode) string {
	return errorMsg[code]
}

func IsSuccess(code ErrorCode) bool {
	return code >= 20000 && code < 30000
}
