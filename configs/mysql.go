package configs

import "time"

type Mysql struct {
	LogMode     bool
	Config      string
	User        string
	Password    string
	Host        string
	Port        int
	Name        string
	Prefix      string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime time.Duration
}
