package initializers

import (
	"fmt"
	"github.com/spf13/viper"
	"gosns/global"
)

func InitConfig() *viper.Viper{
	configFile := "app.toml"
	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: #{err} \n", err))
	}

	if err := v.Unmarshal(&global.CONFIG); err !=nil {
		fmt.Println(err)
	}
	return v
}
