package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sidbusy/weixinmp"
	"github.com/spf13/viper"
)

//go:embed config.toml
var exampleConfig string

func main() {
	CheckConfig()
	InitConfig()

	r := gin.Default()
	r.Any("/wechatmp", WeChatMpHandlers)

	r.Run(viper.GetString("host")) // 监听并在 0.0.0.0:8080 上启动服务
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./data")
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

func CheckConfig() {
	if _, err := os.Stat("./data/config.toml"); err != nil {
		err := os.Mkdir("./data", 0777)
		os.Chmod("./data", 0777)
		if err != nil {
			log.Println(err)
		}

		err = ioutil.WriteFile("./data/config.toml", []byte(exampleConfig), 0666)
		if err != nil {
			log.Println(err)
		}
		println("./data/config.toml 文件已创建并写入内容.")
	} else {
		println("./data/config.toml 文件已存在.")
	}
}
