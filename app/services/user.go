package services

import (
	"gosns/app/models"
	"gosns/global"
	"time"
)

// 初始化用户账号
func CreateAccountWithMobile(mobile string, userID uint) *models.Account{
	account := models.Account{
		Mobile: mobile,
	}
	if userID == 0 {
		now := time.Now()
		user := models.User{
			Status: 1,
			RegTime: now,
			LastLogin: now,
			Birth: time.Unix(0, 0),
		}
		global.DB.Create(&user)
		userID = user.ID
	}
	account.UserID = userID
	global.DB.Create(&account)
	return &account
}