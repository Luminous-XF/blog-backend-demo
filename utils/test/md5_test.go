package test

import (
	"blog-backend/utils"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	password := "abc@123"
	salt := "RyWHw43b1Zgd2ucu"
	str := utils.Md5(password + salt)

	fmt.Println(str)
}
