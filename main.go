package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

var count int

func main(){
	router := gin.Default()

	router.GET("ping", Ping)

	api := router.Group("/api")
	api.POST("/login", UserLogin)
	api.GET("/login", UserLogin)


	router.Run(":8080")
}


func Ping(c *gin.Context){
	count++
	var msg struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Data map[string]int `json:"data"`
	}
	//data :=
	msg.Code = 0
	msg.Msg = "成功"
	msg.Data = map[string]int{"count": count}
	msg.Data["lucky_number"] = rand.Intn(100*100)
	c.JSON(http.StatusOK, msg)
}

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func UserLogin(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err!= nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if user.Username + user.Password != "guye123" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "用户名或者密码错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
}