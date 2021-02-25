package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosns/app/models"
	"gosns/global"
	"net/http"
)

type UserForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


func UserLogin(c *gin.Context) {
	var user UserForm
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if user.Username+user.Password != "guye123" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "用户名或者密码错误"})
		return
	}
	var auser models.User
	//result:=db.Create(&auser)
	//fmt.Println(result)

	global.DB.First(&auser, "age=?", 20)

	fmt.Println(auser)
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功", "user": auser})
}

func UserRegister(c *gin.Context){
	var user UserForm
	if err:=c.ShouldBind(&user); err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var newUser = models.User{}
	newUser.Username = user.Username
	newUser.Password = user.Password
	global.DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}