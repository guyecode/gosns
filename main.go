package main

import (
	"gorm.io/gorm"
	"gosns/global"
	"gosns/initializers"
)



var db *gorm.DB

func main() {
	engine := initializers.InitEngine()

	
	// 加载配置文件
	initializers.InitConfig()
	global.DB = initializers.InitDatabase()
	db, _ :=global.DB.DB()
	defer db.Close()
	engine.Run(":8080")
}
