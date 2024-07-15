package request

type LoginByUsernameAndPasswordRequest struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}
