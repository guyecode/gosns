package initializers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gosns/global"
	"gosns/app/models"
)

func InitDatabase() *gorm.DB {
	mc := global.CONFIG.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mc.User, mc.Password, mc.Host, mc.Port, mc.Name, mc.Config)
	fmt.Println("connetct to mysql use dsn: ", dsn)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		//Logger:                                   getGromLogger(),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mc.Prefix, // 表前缀
			SingularTable: false,      // 使用单数表名
		},
	})

	if err !=nil{
		panic(fmt.Errorf("连接MySQL异常%s:%d", mc.Host, mc.Port))
	}

	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(mc.MaxIdleConns)
	sqldb.SetMaxOpenConns(mc.MaxOpenConns)
	sqldb.SetConnMaxLifetime(mc.MaxLifetime)

	if global.CONFIG.RUN_MODE == gin.DebugMode {
		AutoMigrate(db)
	}
	return db

}


func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&(models.User{}))
	db.AutoMigrate(&(models.Account{}))
	db.AutoMigrate(&(models.Follower{}))
	db.AutoMigrate(&(models.SMSCode{}))
}