package test

import (
	"blog-backend/database"
	"fmt"
	"testing"
)

// go test -v .\database\test\ -run=^TestGetPostById$ -count=1
func TestGetPostById(t *testing.T) {
	post, err := database.GetPostById(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%#v\n", post)
}

// go test -v .\database\test\ -run=^TestGetPostByAuthorId$ -count=1
func TestGetPostByAuthorId(t *testing.T) {
	postList, err := database.GetPostByAuthorId(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for _, post := range postList {
		fmt.Printf("%#v\n", post)
	}
}

func TestUpdatePost(t *testing.T) {
	var post database.Post
	post.Id = 1
	post.Title = "xxx"
	post.Content = "XXX"

	err := database.UpdatePost(&post)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
