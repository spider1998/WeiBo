package app

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug bool `json:"debug" default:"true"`
	//数据库配置
	Mysql string `json:"mysql" default:"root:123456@tcp(192.168.35.193:3306)/weibo?charset=utf8mb4&parseTime=true"` // mysql DSN
	//日志配置
	ConfPath string `json:"conf_path" default:"."` //日志文件路径
	//公告密钥
	AcccessToken string `json:"announce_key" default:"2.005M2CdG0RedLcc972a5c2e8Hz2GmD"`
	//评论限制数
	CommentsCount string `json:"comments_count" default:"200"`
}

//加载配置
func LoadConfig() (Config, error) {
	godotenv.Load()
	var config Config
	err := envconfig.Process("", &config)
	return config, err
}
