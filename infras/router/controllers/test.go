package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (controller *Controllers) GetTestData(ctx *gin.Context) {
	fmt.Println("GetTestData")
}
