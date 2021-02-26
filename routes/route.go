package routes

import (
	"github.com/gin-gonic/gin"
	"gosns/global"
)

func InitRoute() *gin.Engine{
	gin.SetMode(global.CONFIG.RUN_MODE)
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	loadRoute(engine)
	return engine

}

func loadRoute(engine *gin.Engine){
	loadApiRoutes(engine)
	loadTestRoutes(engine)
}
