package models

type CreatePostModel struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	UserId int64  `json:"user_id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type GetPostByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
