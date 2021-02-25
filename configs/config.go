package configs


type Config struct {
	DEBUG bool
	Mysql Mysql `mapstructure:"mysql"`
}