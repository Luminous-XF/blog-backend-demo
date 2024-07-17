package test

import (
	"blog-backend/global"
	"blog-backend/middleware"
	"blog-backend/model/token"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestGenToken(T *testing.T) {
	uid, _ := uuid.Parse("2c936d7e-4247-11ef-a62c-20906f8b3d78")

	j := &middleware.JWT{
		SigningKey: []byte(global.CONFIG.JWTConfig.SigningKey),
	}

	claims := token.CustomClaims{
		UUID: uid,
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
		T.Error(err.Error())
		T.Fail()
	}
	fmt.Printf("tokenStr: %s\n", tokenStr)

	c, err := j.ParseToken(tokenStr)
	if err != nil {
		T.Error(err.Error())
		T.Fail()
	}
	fmt.Printf("claims-UUID: %#v\n", c.UUID.String())
	fmt.Printf("claims: %#v\n", c.RegisteredClaims)
}
