package service

import (
	"blog-backend/common/error_code"
	"blog-backend/database"
	"blog-backend/model/request"
	"blog-backend/model/response"
)

func GetPostList(info request.PageInfoRequest) (postResponseList []*response.PostResponse, code error_code.ErrorCode) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	postList, err := database.GetPostList(offset, limit)
	if err != nil {
		return nil, error_code.QueryPostListFail
	}

	for _, post := range postList {
		postResponseList = append(postResponseList, &response.PostResponse{
			UUID:           post.UUID,
			Title:          post.Title,
			Excerpt:        post.Excerpt,
			Content:        post.Content,
			Type:           post.Type,
			CommentCount:   post.CommentCount,
			Score:          post.Score,
			Status:         post.Status,
			Likes:          post.Likes,
			PageViews:      post.PageViews,
			HeaderImageUrl: post.HeaderImageUrl,
			CreateTime:     post.CreateTime,
			UpdateTime:     post.UpdateTime,
		})
	}

	return postResponseList, error_code.SUCCESS
}
