package test

import (
	"blog-backend/database"
	"blog-backend/model"
	"blog-backend/utils"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

// go test -v .\database\test\ -run=^TestGetUserByUsername$ -count=1
func TestGetUserByUsername(t *testing.T) {
	user, err := database.GetUserByUsername("19380120319")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%#v\n", user)
}

// go test -v .\database\test\ -run=^TestGetUserByID$ -count=1
func TestGetUserByID(t *testing.T) {
	user, err := database.GetUserByID(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%#v\n", user)
}

// go test -v .\database\test\ -run=^TestGetUserByUUID$ -count=1
func TestGetUserByUUID(t *testing.T) {
	user, err := database.GetUserByUUID("c7ac28ba-3fdd-11ef-a62c-20906f8b3d78")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%#v\n", user)
}

// go test -v .\database\test\ -run=^TestCreateUser$ -count=1
func TestCreateUser(t *testing.T) {
	uid := uuid.New()
	salt := utils.MakeStr(16, utils.DigitAlphaPunct)
	var user = model.User{
		UUID:     uid,
		Username: "IU",
		Password: utils.Md5("abc@123" + salt),
		Salt:     salt,
		Email:    utils.MakeStr(10, utils.Alpha) + "@gmail.com",
	}

	if err := database.CreateUser(&user); err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%#v\n", user)
}
