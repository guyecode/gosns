package routes

import (
	"github.com/gin-gonic/gin"
	"gosns/app/controllers"
	"gosns/app/middlewares"

)

func loadApiRoutes(engine *gin.Engine){
	api := engine.Group("/api")
	api.Use(middlewares.Auth())
	api.POST("/login", controllers.UserLogin)
	api.GET("/register", controllers.UserRegister)
}



