package routes

import (
	"github.com/gin-gonic/gin"
	"gosns/app/controllers"
)

func loadOpenApiRoutes(engine *gin.Engine){
	api := engine.Group("/api")
	api.GET("/login", controllers.UserLogin)
	api.GET("/register", controllers.UserRegister)
	api.POST("/sms", controllers.SendCode)

}



