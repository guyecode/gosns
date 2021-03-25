package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosns/app/models"
	"gosns/app/services"
	"gosns/utils"
	"net/http"
)


type MobileLoginForm struct {
	Mobile string `form:"mobile" binding:"required,mobile_number"`
	Code string `form:"code" binding:"required,numeric,len=6"`
}


func LoginWithMobile(c *gin.Context) {
	var form MobileLoginForm
	if err := c.ShouldBind(&form); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		return
	}
	if !services.CheckSMSCode(form.Mobile, form.Code) {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "验证失败"})
		return
	}

	account := models.GetAccountByMobile(form.Mobile)
	fmt.Println(account)
	if utils.IsNil(account) {
		account = services.CreateAccountWithMobile(form.Mobile, 0)
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "登录成功"})
}