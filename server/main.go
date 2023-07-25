package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sidbusy/weixinmp"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()

	r := gin.Default()
	r.Any("/wechatmp", WeChatMpHandlers)

	r.Run(viper.GetString("host")) // 监听并在 0.0.0.0:8080 上启动服务
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	log.Println("Init Config")
	// Initialize WechatMP Configuration
	mpConfig := viper.GetStringMap("wechat")
	WeChatMp = weixinmp.New(
		mpConfig["token"].(string),
		mpConfig["appid"].(string),
		mpConfig["secret"].(string),
	)
}
