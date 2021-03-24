package models

import (
	"fmt"
	"gosns/global"
	"time"
)

// 短信验证码
type SMSCode struct {
	Model
	Mobile string `gorm:"index"`
	Code   string
	Device string
	IP     string
}

func GetSMSByMobile(mobile string) (sms SMSCode) {
	global.DB.Where("mobile=?", mobile).Last(&sms)
	return sms
}


func CountSMSToday(field string, value string) (count int64) {
	var sms SMSCode
	now := time.Now()
	todayZeroClock := time.Date(now.Year(), now.Month(), now.Day(), 0, 0,0, 0, now.Location())
	sql := fmt.Sprintf("%s=? and created_at>?", field)
	global.DB.Model(&sms).Where(sql, value, todayZeroClock).Count(&count)
	return
}