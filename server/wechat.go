package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sidbusy/weixinmp"
	"github.com/spf13/viper"
)

var WeChatMp *weixinmp.Weixinmp

func WeChatMpHandlers(c *gin.Context) {
	if !WeChatMp.Request.IsValid(c.Writer, c.Request) {
		log.Println("Not a valid request")
		return
	}
	if !FindAOpenIDExist(WeChatMp.Request.FromUserName) {
		log.Println("陌生人:", WeChatMp.Request.FromUserName)
		WeChatMp.ReplyTextMsg(c.Writer, viper.GetString("wechat_r_unkown_user"))
		return
	}
	openAPIUrl := FindByWechatOpenID(WeChatMp.Request.FromUserName, "memos_open_api")

	r_str := viper.GetString("wechat_default_r_text")
	if r_str == "" {
		r_str = "📩 已保存"
	}
	var err error
	var content string
	var visibility = "PRIVATE"
	switch WeChatMp.Request.MsgType {
	case weixinmp.MsgTypeText: // 文字消息
		r_str = viper.GetString("wechat_r_text")
		content, visibility = PublicText(WeChatMp.Request.Content)
		err = CreateMemo(openAPIUrl, content, visibility, []int{})
	case weixinmp.MsgTypeImage: // 图片消息
		r_str = viper.GetString("wechat_r_image")
		var id int
		id, err = CreateResourceByLink(openAPIUrl, WeChatMp.Request.PicUrl)
		if err != nil {
			err = CreateMemo(openAPIUrl, "", visibility, []int{id})
		}
	case weixinmp.MsgTypeVoice: // 语言消息
		r_str = viper.GetString("wechat_r_voice")
		if WeChatMp.Request.Recognition != "" {
			content, visibility = PublicText(WeChatMp.Request.Recognition)
			err = CreateMemo(openAPIUrl, content, visibility, []int{})
		} else {
			r_str = "没有识别到文字"
		}
	case weixinmp.MsgTypeLocation: // 位置消息
		r_str = viper.GetString("wechat_r_location")
		content = fmt.Sprintf("位置信息: 位置 %s <br>经纬度( %f , %f )", WeChatMp.Request.Label, WeChatMp.Request.LocationX, WeChatMp.Request.LocationY)
		err = CreateMemo(openAPIUrl, content, visibility, []int{})
	case weixinmp.MsgTypeLink: // 链接消息
		r_str = viper.GetString("wechat_r_link")
		content = fmt.Sprintf("[%s](%s)\n%s...", WeChatMp.Request.Title, WeChatMp.Request.Url, WeChatMp.Request.Description)
		err = CreateMemo(openAPIUrl, content, visibility, []int{})
	case weixinmp.MsgTypeVideo:
		r_str = viper.GetString("wechat_r_video")
	default:
		r_str = viper.GetString("wechat_r_unknown")
	}
	if err != nil {
		log.Println(err)
		r_str = "Error"
	}
	// 替换变量
	r_str = ReplaceVariables(WeChatMp.Request.FromUserName, r_str)
	r_str = strings.ReplaceAll(r_str, "${content}", content)
	r_str = strings.ReplaceAll(r_str, "${visibility}", visibility)
	WeChatMp.ReplyTextMsg(c.Writer, r_str)
}

// input := "这是一个示例字符串，${变量1} 和 ${变量2} ${openid} ${memos_open_api}需要替换为实际的值。"
func ReplaceVariables(wechatopenid string, input string) string {
	re := regexp.MustCompile(`\${([^}]+)}`)
	output := re.ReplaceAllStringFunc(input, func(match string) string {
		variable := re.FindStringSubmatch(match)[1]
		return FindByWechatOpenID(wechatopenid, variable)
	})
	return output
}

// content PRIVATE/PUBLIC
func PublicText(input string) (string, string) {
	if strings.HasPrefix(input, "公开发布") {
		result := strings.TrimPrefix(input, "公开发布")
		result = strings.TrimLeft(result, "，。：")
		return result, "PUBLIC"
	} else {
		return input, "PRIVATE"
	}
}
