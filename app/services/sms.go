package services

import (
	"fmt"
	"gosns/app/models"
	"time"
)

var (
	// 短信发送间隔
	SMS_SEND_INTERVAL float64 = 60
	// 同一手机号每天可以发送的短信数量
	SMS_MOBILE_PER_DAY int64 = 5
	// 同一台设备每天可以发送的短信数量
	SMS_DEVICE_PER_DAY int64 = 10
	// 同一IP每天可以发送的短信数量
	SMS_IP_PER_DAY int64 = 20
	// 短信签名，运营商要求，必需要带的。
	SMS_SIGNATURE = "测试"
	// 验证码过期时长（秒）
	SMS_CODE_EXPIRE_DURATION float64 = 60 * 5
)



// 检查是否符合发送条件
func LimitSMSSend(mobile string, device string, ip string) (ok bool, reason string){
	sms := models.GetSMSByMobile(mobile)
	now := time.Now()
	duration := now.Sub(sms.CreatedAt)
	if duration.Seconds() < SMS_SEND_INTERVAL {
		return false, fmt.Sprintf("同一手机号发送间隔不得小于%.0f秒", SMS_SEND_INTERVAL)
	}

	if models.CountSMSToday("mobile", mobile) >= SMS_MOBILE_PER_DAY {
		return false, "该手机号已达到当天最大发送短信次数"
	}
	if models.CountSMSToday("device", device) >= SMS_DEVICE_PER_DAY{
		return false, "SMS send over capacity"
	}
	if models.CountSMSToday("ip", ip) >= SMS_IP_PER_DAY{
		return false, "SMS send over capacity."
	}
	return true, ""
}

// 发送短信
func SendSMS(mobile string, content string) {
	fmt.Printf("sending SMS to %s: %s\n", mobile, content)
}

// 发送短信验证码
func SendCodeSMS(mobile string, code string) {
	content := fmt.Sprintf("【%s】您的验证码是 %s，此验证码%.0f分钟内有效，请勿透露给其他人。",
		SMS_SIGNATURE, code, SMS_CODE_EXPIRE_DURATION / 60)
	SendSMS(mobile, content)
}

// 检查验证码是否有效并正确
func CheckSMSCode(mobile string, code string) bool {
	sms := models.GetSMSByMobile(mobile)
	if sms == (models.SMSCode{}) || sms.Code != code  {
		return false
	}
	now := time.Now()
	duration := now.Sub(sms.CreatedAt)
	if duration.Seconds() > SMS_CODE_EXPIRE_DURATION {
		return false
	}
	return true
}