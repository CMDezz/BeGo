package controllers

import "github.com/gin-gonic/gin"

type IUserControllers interface {
	GetAllUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUserById(ctx *gin.Context)
}

func (controller *Controllers) GetAllUsers(ctx *gin.Context) {

}

func (controller *Controllers) GetUserById(ctx *gin.Context) {

}

func (controller *Controllers) CreateUser(ctx *gin.Context) {

}

func (controller *Controllers) UpdateUser(ctx *gin.Context) {

}

func (controller *Controllers) DeleteUserById(ctx *gin.Context) {

}
