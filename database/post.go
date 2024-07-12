package database

import (
	"blog-backend/global"
	"blog-backend/model"
)

// GetPostById 通过帖子 id 查询帖子
func GetPostById(id int) (post *model.Post, err error) {
	err = global.GDB.Where("id = ?", id).First(&post).Error
	return post, err
}

// GetPostByAuthorId 根据 user_id 查询该用户发布的帖子列表
func GetPostByAuthorId(userId int) (postList []*model.Post, err error) {
	err = global.GDB.Where("author_id = ?", userId).Find(&postList).Error
	return postList, err
}

// UpdatePost 修改帖子标题和内容
func UpdatePost(post *model.Post) (newPost *model.Post, err error) {
	err = global.GDB.Where("id = ?", post.ID).First(&newPost).Updates(model.Post{
		Title:   post.Title,
		Content: post.Content,
	}).Error
	return newPost, err
}
