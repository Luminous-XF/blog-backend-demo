package request

type LoginByUsernameAndPasswordRequest struct {
	Username string `json:"username" binding:"required,min=3,max=16,username-charset"`
	Password string `json:"password" binding:"required,min=8,max=16,password-charset"`
}
