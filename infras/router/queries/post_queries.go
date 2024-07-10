package queries

import (
	"BeGo/constants"
	"BeGo/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type IPostQueries interface {
	QueryGetAllPost(ctx *gin.Context) (*[]models.Post, error)
	QueryGetPostById(ctx *gin.Context, id int64) (*models.Post, error)
	QueryCreatePost(ctx *gin.Context, post models.Post) (*models.Post, error)
}

func (queries *Queries) QueryGetAllPost(ctx *gin.Context) (*[]models.Post, error) {
	query := fmt.Sprintf("SELECT * FROM %v",
		constants.TABLE_POSTS,
	)

	var res []models.Post

	err := queries.db.SelectContext(ctx, &res, query)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) QueryGetPostById(ctx *gin.Context, id int64) (*models.Post, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE id=%d",
		constants.TABLE_POSTS, id,
	)
	var res models.Post

	err := queries.db.GetContext(ctx, &res, query)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (queries *Queries) QueryCreatePost(ctx *gin.Context, post models.Post) (*models.Post, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, body, user_id,status ) VALUES ($1, $2, $3,$4) RETURNING *",
		constants.TABLE_POSTS,
	)

	res := models.Post{}
	err := queries.db.QueryRowxContext(ctx, query, post.Title, post.Body, post.UserId, post.Status).StructScan(&res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
