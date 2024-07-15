package test

import (
	"blog-backend/database"
	"blog-backend/model"
	"blog-backend/model/request"
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

// go test -v .\database\test\ -run=^TestUpdatePost$ -count=1
func TestUpdatePost(t *testing.T) {
	var post model.Post
	post.ID = 1
	post.Title = "111"
	post.Content = "222"
	post.Excerpt = "333"

	newPost, err := database.UpdatePost(&post)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%#v\n", newPost)
}

// go test -v .\database\test\ -run=^TestGetPostListByCreateTime$ -count=1
func TestGetPostListByCreateTime(t *testing.T) {
	pageInfo := request.PageInfoRequest{
		Page:     1,
		PageSize: 1,
	}
	postList, err := database.GetPostList(pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for _, post := range postList {
		fmt.Printf("%#v\n", post)
	}
}
