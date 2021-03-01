package routes

import (
	"github.com/gin-gonic/gin"
	"gosns/app/controllers"
	"gosns/app/middlewares"
)

func loadApiRoutes(engine *gin.Engine){
	api := engine.Group("/api")
	api.Use(middlewares.Auth())
	api.GET("/login", controllers.UserLogin)
	api.GET("/register", controllers.UserRegister)
	api.GET("/user/profile/:id", controllers.UserProfile)
}



