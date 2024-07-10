package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	// ErrorCode    int //tạm thời theo http code, nếu implement app error code thì sử dụng ở đây
	Message      string
	ErrorMessage string
}

type DataResponse struct {
	// Code    int //tạm thời theo httpcode, nếu implement app success code thì sử dụng ở đây
	Message string
	Data    any
}

func ResponseBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, GinError(err))
}

func ResponseInternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, GinError(err))
}

func ResponseUnauthorized(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, GinError(err))
}

func ResponseOk(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, DataResponse{
		Message: "Successfully!",
		Data:    data,
	})
}

func GinError(err error) ErrorResponse {
	return ErrorResponse{
		Message:      "Oops! Something wrong",
		ErrorMessage: err.Error(),
	}
}
