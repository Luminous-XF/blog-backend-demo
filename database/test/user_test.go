package test

import (
	"blog-backend/database"
	"fmt"
	"testing"
)

func TestGetUserByUsername(t *testing.T) {
	user, err := database.GetUserByUsername("193801203191")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%#v\n", user)
}
