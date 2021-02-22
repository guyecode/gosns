package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gosns/configs"
)

var (
	CONFIG configs.Config
	VIPER *viper.Viper
	DB *gorm.DB
)