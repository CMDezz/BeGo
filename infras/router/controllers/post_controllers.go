package controllers

import (
	"BeGo/infras/response"
	"BeGo/models"

	"github.com/gin-gonic/gin"
)

type IPostControllers interface {
	GetAllPosts(ctx *gin.Context)
	GetPostById(ctc *gin.Context)
	CreatePost(ctx *gin.Context)
}

func (controller *Controllers) GetAllPosts(ctx *gin.Context) {
	res, err := controller.Queries.QueryGetAllPost(ctx)

	if err != nil {
		response.ResponseInternalServerError(ctx, err)
		return
	}
	response.ResponseOk(ctx, res)
}

func (controller *Controllers) GetPostById(ctx *gin.Context) {
	var req models.GetPostByIdRequest

	//validate req uri
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ResponseBadRequest(ctx, err)
		return
	}

	res, err := controller.Queries.QueryGetPostById(ctx, req.ID)

	if err != nil {
		response.ResponseInternalServerError(ctx, err)
		return
	}
	response.ResponseOk(ctx, res)
}

func (controller *Controllers) CreatePost(ctx *gin.Context) {
	var req models.CreatePostModel

	//validate req json
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ResponseBadRequest(ctx, err)
		return
	}

	post := models.Post{
		Title:  req.Title,
		Body:   req.Body,
		UserId: req.UserId,
		Status: req.Status,
	}

	res, err := controller.Queries.QueryCreatePost(ctx, post)

	if err != nil {
		response.ResponseInternalServerError(ctx, err)
		return
	}
	response.ResponseOk(ctx, res)
}
