package controllers

import "github.com/gin-gonic/gin"

type Controllers struct {
}

type IController interface {
	GetTestData(ctx *gin.Context)
}
