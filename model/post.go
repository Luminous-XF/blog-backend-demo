package model

import "time"

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
	Model
	UUID           string    `gorm:"column:uuid"`
	AuthorId       int64     `gorm:"author_id"`
	Title          string    `gorm:"title"`
	Excerpt        string    `gorm:"column:excerpt"`
	Content        string    `gorm:"content"`
	Type           Type      `gorm:"type"`
	CommentCount   int64     `gorm:"comment_count"`
	Score          float64   `gorm:"score"`
	Status         Status    `gorm:"status"`
	Likes          int64     `gorm:"likes"`
	PageViews      int64     `gorm:"page_views"`
	HeaderImageUrl string    `gorm:"header_image_url"`
	CreateTime     time.Time `gorm:"create_time"`
	UpdateTime     time.Time `gorm:"update_time"`
}

func (Post) TableName() string {
	return "post"
}
