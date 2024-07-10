package apis

import (
	"BeGo/constants"
	"BeGo/infras/router/controllers"
	"BeGo/infras/router/queries"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	router      *gin.Engine
	Controllers controllers.IControllers
	// Queries     queries.IQueries
}

func InitServer(db *sqlx.DB) (*Server, error) {
	queries := queries.NewQueries(db)
	controller := controllers.NewControllers(queries)

	server := &Server{
		Controllers: controller,
	}

	server.InitRouter()

	return server, nil
}

func (server *Server) InitRouter() {
	//Init Gin
	router := gin.Default()

	//Sample apis
	router.GET(constants.API_V1+"/GetAllPost", server.Controllers.GetAllPosts)
	router.POST(constants.API_V1+"/CreatePost", server.Controllers.CreatePost)
	router.GET(constants.API_V1+"/GetPostById/:id", server.Controllers.GetPostById)

	server.router = router
}

func (server *Server) Start() {
	server.router.Run()
}
