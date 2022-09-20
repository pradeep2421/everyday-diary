package utils

import "github.com/spf13/viper"
type DbConfig struct {
	DbDriver      string `mapstructure:"driver"`
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	User string `mapstructure:"user"`
}
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Address string `mapstructure:"address"`
}
type Config struct {
	Db DbConfig `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
}
var vp *viper.Viper
func LoadConfig(path string) (config Config, err error) {
	vp = viper.New()
	vp.AddConfigPath(".")
	vp.AddConfigPath("./utils")
	vp.AddConfigPath(path)
	vp.SetConfigName("config")
	vp.SetConfigType("json")

	err = vp.ReadInConfig()
	if err != nil {
		return
	}

	err = vp.Unmarshal(&config)
	return
}


// DB_HOST=localhost
// DB_USER=root
// DB_PASSWORD=secret
// DB_PORT=3306
// DB_DATABASE_NAME= diary
// DB_DRIVER=mysql
// SERVER_ADDRESS=0.0.0.0:8080
