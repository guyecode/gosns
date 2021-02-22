package routes

import "github.com/gin-gonic/gin"

func InitRoute() *gin.Engine{
	engine := gin.New()
	loadRoute(engine)
	return engine

}

func loadRoute(engine *gin.Engine){
	loadApiRoutes(engine)
	loadTestRoutes(engine)
}
