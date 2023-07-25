# Memos-Wechat-Server

For memos 

！！！未发布，待测试！！！

配置文件，可使用变量

eg: `wechat_r_unkown_user = "陌生人: ${openid} ${some}"`

eg: `wechat_r_default = "📩 已保存 ${content} ${visibility}"`

微信测试号 URL 路径: `/wechatmp`

关注后获取 openid , 配置文件 memos_open_api 填写对应 memos openapi url.

可以云函数部署，教程暂时参考 https://www.ftls.xyz/posts/obcsapi-fc-simple/ 云函数部分

Docker kkbt/memos-wechatmp-server 待测试

```toml
host = "0.0.0.0:8080"


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