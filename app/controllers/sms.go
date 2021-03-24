package controllers

import (
	"github.com/gin-gonic/gin"
	"gosns/app/models"
	"gosns/app/services"
	"gosns/global"
	"gosns/utils"
	"net/http"
)

type SendCodeForm struct {
	DeviceID string `form:"deviceid" json:"deviceID" binding:"required"`
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile_number"`
}

// @Summary 发送手机验证码
// @Description  用户注册、找回密码时需要用到
// @Tags 手机短信
// @Produce  json
// @Param deviceid header string false "设备ID"
// @Param mobile formData string true "手机号"
// @Failure 400 {object} utils.SKELETON "参数错误"
// @Failure 1 {object} string "发送达到上限"
// @Success 0 {object} UserSwagger "发送成功"
// @Router /api/sms [POST]
func SendCode(c *gin.Context) {
	var form SendCodeForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	ok, reason := services.LimitSmsSend(form.Mobile, form.DeviceID, c.ClientIP())
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": reason})
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
