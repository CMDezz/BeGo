package queries

import (
	"github.com/gin-gonic/gin"
)

type IUserQueries interface {
	QueryGetAllUser(ctx *gin.Context)
}

func (queries *Queries) QueryGetAllUser(ctx *gin.Context) {

}
