package app

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug bool `json:"debug" default:"true"`
	//数据库配置
	Mysql     string `json:"mysql" default:"root:123456@tcp(192.168.35.193:3306)/weibo?charset=utf8mb4&parseTime=true"` // mysql DSN
	Redis     string `json:"redis" default:"192.168.35.193:6379"`
	LikeRedis string `json:"like_redis" default:"goblog"`
	//日志配置
	ConfPath string `json:"conf_path" default:"."` //日志文件路径
	//公告密钥
	AcccessToken string `json:"announce_key" default:"2.005M2CdGDsxwmC737c94b50c0t5yfU"`
}

//加载配置
func LoadConfig() (Config, error) {
	godotenv.Load()
	var config Config
	err := envconfig.Process("", &config)
	return config, err
}
