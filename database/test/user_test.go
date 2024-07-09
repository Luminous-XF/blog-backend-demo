package test

import (
	"blog-backend/database"
	"blog-backend/util"
	"fmt"
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

// go test -v .\database\test\ -run=^TestCreateUser$ -count=1
func TestCreateUser(t *testing.T) {
	database.CreateUser("IU", util.GenerateSalt(8)+"@qq.com", "abc@123")
}

func TestDeleteUserByUsername(t *testing.T) {
	if err := database.DeleteUserByUsername("IU"); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
