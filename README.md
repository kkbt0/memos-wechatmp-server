# Memos-WechatMP-Server

用于转发微信测试号/公众号消息到 memos 
For memos https://github.com/usememos/memos 

## 微信测试号

微信测试号 URL 路径: `/wechatmp`

关注后获取 openid , 配置文件 memos_open_api 填写对应 memos openapi url. token ，appid 等等从微信测试号管理页面获设置和获取。

可以云函数部署，教程暂时参考 https://www.ftls.xyz/posts/obcsapi-fc-simple/ 云函数部分

Docker kkbt/memos-wechatmp-server 完成测试 最新版本 v202307281434

```bash
docker run -p 8905:8905 -v /app/memos-wechatmp-server/data/:/app/data/ kkbt/memos-wechatmp-server:v202307281434
```

## 配置

配置文件，可使用变量

eg: `wechat_r_unkown_user = "陌生人: ${openid} ${some}"`

eg: `wechat_r_default = "📩 已保存 ${content} ${visibility}"`

使用举例：`wechat_r_default = "📩 已保存 ${front_end_url}"` ，然后配置变量  front_end_url 为 Memos 前端链接。

## 使用

可发送图片，语音，文字，链接等。语音，文字发送 以 公开发布 开头，则为公开发布语音或文字。

## 配置示例

```toml
host = "0.0.0.0:8905"


users = [
    { openid = "xxx", memos_open_api = "https://demo.usememos.com/api/v1/memo?openId=B0A20B12622CD78F448856BD67F1EF7A" },
    { openid = "kkbt", memos_open_api = "http://192.168.0.107:5230/api/memo?openId=b081854f-77b7-4dac-9ced-626c37d39edc" , some:"sometext"},
]

wechat_r_unkown_user = "陌生人: ${openid}"
wechat_r_default = "📩 已保存"
wechat_r_text = "📩 已保存文字"
wechat_r_image = "📩 已保存图片"
wechat_r_voice = "📩 已保存语音"
wechat_r_location = "📩 已保存位置"
wechat_r_link = "📩 已保存链接"
wechat_r_video ="不支持视频消息"
wechat_r_unknown = "未知类型消息"

[wechat]
token = "wechat_token"
appid = "wechat_appid"
secret = "wechat_secret"
```
