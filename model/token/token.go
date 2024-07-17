package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type CustomClaims struct {
	UUID       uuid.UUID `json:"uuid"`
	Username   string    `json:"username"`
	BufferTime int64     `json:"buffer_time"`
	jwt.RegisteredClaims
}
