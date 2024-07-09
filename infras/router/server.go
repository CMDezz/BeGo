package apis

import (
	"BeGo/infras/router/controllers"
	"BeGo/infras/router/queries"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	Controllers controllers.IController
	Queries     queries.Queries
}

func InitServer() (*Server, error) {

	server := &Server{}
	server.InitRouter()
	return server, nil
}

func (server *Server) InitRouter() {
	//Init Gin
	router := gin.Default()

	//Sample apis
	router.GET("/test")

	server.router = router
}

func (server *Server) Start() {
	server.router.Run()
}
