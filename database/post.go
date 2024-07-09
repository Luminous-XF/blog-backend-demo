package database

import (
	"blog-backend/util"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Type 帖子类型枚举
type Type string

const (
	REGULAR Type = "REGULAR"
	HOT     Type = "HOT"
	TOP     Type = "HOT"
)

// Status 帖子状态枚举
type Status string

const (
	NORMAL  Status = "NORMAL"
	REVIEW  Status = "REVIEW"
	BLOCKED Status = "BLOCKED"
)

type Post struct {
	Uuid         string    `gorm:"column:uuid"`
	Id           int       `gorm:"id;AUTO_INCREMENT;primary_key"`
	AuthorId     int       `gorm:"author_id"`
	Title        string    `gorm:"title"`
	Content      string    `gorm:"content"`
	Type         Type      `gorm:"type"`
	CommentCount int64     `gorm:"comment_count"`
	Score        float64   `gorm:"score"`
	Status       Status    `gorm:"status"`
	Likes        int64     `gorm:"likes"`
	PageViews    int64     `gorm:"page_views"`
	CreateTime   time.Time `gorm:"create_time"`
	UpdateTime   time.Time `gorm:"update_time"`

	RowVersion    int       `gorm:"column:row_version"`
	RowCreateTime time.Time `gorm:"column:row_create_time"`
	RowUpdateTime time.Time `gorm:"column:row_update_time"`
	RowIsDeleted  int       `gorm:"column:row_is_deleted"`
}

func (Post) TableName() string {
	return "post"
}

var (
	allPostField = util.GetGormFields(Post{})
)

// GetPostById 通过帖子 id 查询帖子
func GetPostById(id int) (*Post, error) {
	db := GetBlogDBConnection()

	var post Post
	if err := db.Select(allPostField).Where("id = ?", id).First(&post).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			util.Logger.Errorf("get password of user '%d' failed: '%s'", id, err)
		}
		return nil, err
	}

	return &post, nil
}

// GetPostByAuthorId 根据 user_id 查询该用户发布的帖子列表
func GetPostByAuthorId(userId int) ([]*Post, error) {
	db := GetBlogDBConnection()

	var posts []*Post
	if err := db.Select(allPostField).Where("author_id = ?", userId).Find(&posts).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			util.Logger.Errorf("get post by user '%d' failed: '%s'", userId, err)
		}
		return nil, err
	}

	return posts, nil
}

func UpdatePost(post *Post) error {
	if post.Id <= 0 {
		return fmt.Errorf("could not update blog of id = '%d'", post.Id)
	}

	if len(post.Title) == 0 || len(post.Content) == 0 {
		return fmt.Errorf("could not set blog title or content to empty")
	}

	db := GetBlogDBConnection()
	err := db.Model(Post{}).Where("id = ?", post.Id).Updates(map[string]any{
		"title":   post.Title,
		"content": post.Content,
	}).Error

	return err
}
