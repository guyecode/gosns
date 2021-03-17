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
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedAt  int `json:"deleted_at" gorm:"-"`
}

type User struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

type UserSwagger struct {
	Lists []*User
	Total int
}

type MobileLoginForm struct {
	Mobile string `form:"mobile"  binding:"required"`
	Code string `form:"code" binding:"required"`
}


func LoginWithMobile(c *gin.Context) {
	var form MobileLoginForm
	if err := c.ShouldBind(&form); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

// @Summary 用户登录
// @Description  获取用户列表
// @Tags 用户信息
// @Produce  json
// @Failure 400 {object} utils.SKELETON "参数错误"
// @Failure 20001 {object} utils.SKELETON "Token鉴权失败"
// @Failure 20002 {object} utils.SKELETON "Token已超时"
// @Failure 20004 {object} utils.SKELETON "Token错误"
// @Failure 20005 {object} utils.SKELETON "Token参数不能为空"
// @Success 0 {object} UserSwagger "查询成功"
// @Router /api/login [get]
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
	// result:=db.Create(&auser)
	// fmt.Println(result)

	global.DB.First(&auser, "age=?", 20)

	fmt.Println(auser)
	c.SetCookie("token", "abcd", 3600, "/", "127.0.0.1", false, true)
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功", "user": auser})
}


// @Summary 用户注册
// @Description  使用用户名密码注册成为新用户
// @Description 用户名不能为空，密码不能为空
// @Tags 用户信息
// @Produce  json
// @Param username query string false "用户名"
// @Param password query string true "密码"
// @Failure 400 {object} utils.SKELETON "参数错误"
// @Failure 20001 {object} utils.SKELETON "Token鉴权失败"
// @Failure 20002 {object} utils.SKELETON "Token已超时"
// @Failure 20004 {object} utils.SKELETON "Token错误"
// @Failure 20005 {object} utils.SKELETON "Token参数不能为空"
// @Success 0 {object} UserSwagger "查询成功"
// @Router /api/register [get]
func UserRegister(c *gin.Context) {
	var user UserForm
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	var newUser = models.Account{}
	newUser.Username = user.Username
	newUser.Password = user.Password
	global.DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func UserProfile(c *gin.Context) {
	id := c.Param("id")
	var user = models.User{}
	global.DB.First(&user, id)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
