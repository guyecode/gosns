package controllers

import (
	"github.com/gin-gonic/gin"
	"gosns/app/models"
	"gosns/global"
	"gosns/utils"
	"net/http"
)

type SendCodeForm struct {
	DeviceID string `form:"deviceid" json:"deviceID" binding:"required"`
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
}

// @Summary 发送手机验证码
// @Description  用户注册、找回密码时需要用到
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
func SendCode(c *gin.Context) {
	var form SendCodeForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	sms := models.SMSCode{
		Mobile: form.Mobile,
		Device: form.DeviceID,
		IP: c.ClientIP(),
		Code: utils.RandNumbers(6),
	}
	global.DB.Create(&sms)
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
