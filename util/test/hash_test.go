package test

import (
	"blog-backend/util"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	password := "abc@123"
	salt := "RyWHw43b1Zgd2ucu"
	str := util.Md5(password + salt)

	fmt.Println(str)
}
