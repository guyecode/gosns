package routes

import (
	"github.com/gin-gonic/gin"
	"gosns/app/controllers"
)

func loadApiRoutes(engine *gin.Engine){
	api := engine.Group("/api")
	api.POST("/login", controllers.UserLogin)
	api.GET("/register", controllers.UserRegister)
}



