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
)



// 检查是否符合发送条件
func LimitSmsSend(mobile string, device string, ip string) (ok bool, reason string){
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