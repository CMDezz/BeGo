package controllers

import "github.com/gin-gonic/gin"

type Controllers struct {
}

type IControllers interface {
	GetTestData(ctx *gin.Context)
}

func NewControllers() IControllers {
	return &Controllers{}
}
