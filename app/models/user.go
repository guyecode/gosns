package models

import (
	"time"
)

type Account struct {
	Model
	Username string	`json:"-"`
	Password string `json:"-"`
	Email string `json:"email"`
	Mobile string `json:""`
	UserID uint `json:"userid"`
	WechatID string
	WeiboID string
	QQID string
}

type User struct {
	Model
	Nickname string `json:"name"`
	Age int	`json:"age"`
	Gender uint `json:"gender"`
	Avatar string `json:"avatar"`
	Region string `json:"region"`
	Birth time.Time
	Intro string `json:"intro"`
	FansNum int `json:"fans_num"`
	FollowNum int `json:"follow_num"`
	Popularity int `json:"popularity"`
	RegTime time.Time `json:"reg_time"`
	LastLogin time.Time `json:"last_login"`
	Status int `json:"status"`
}

// 用户关系
type Follower struct {
	Model
	From uint `gorm:"index"`
	To uint `gorm:"index"`
	Status int
	Group string
}

