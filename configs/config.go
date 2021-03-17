package configs


type Config struct {
	RUN_MODE string `mapstructure:"RUN_MODE"`
	Mysql Mysql `mapstructure:"mysql"`
	Redis Redis `mapstructure:"redis"`
}