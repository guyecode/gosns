package initializers

import (
	"github.com/gin-gonic/gin"
	"gosns/routes"
)

func InitEngine() *gin.Engine{
	return routes.InitRoute()
}
