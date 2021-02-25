package routes

import (
	"github.com/gin-gonic/gin"
	"gosns/global"
)

func InitRoute() *gin.Engine{
	gin.SetMode(global.CONFIG.RUN_MODE)
	engine := gin.Default()
	loadRoute(engine)
	return engine

}

func loadRoute(engine *gin.Engine){
	loadApiRoutes(engine)
	loadTestRoutes(engine)
}
