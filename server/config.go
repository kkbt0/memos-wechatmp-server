package main

import (
	"github.com/spf13/viper"
)

type User struct {
	OpenID      string `toml:"openid"`
	MemosAPIURL string `toml:"memos_open_api"`
}

type Config struct {
	Host  string `toml:"host"`
	Users []User `toml:"users"`
}

// 根据 openid 查到对应 key
func FindByWechatOpenID(wechatopenid string, index string) string {
	usersList, ok := viper.Get("users").([]interface{})
	if !ok {
		return ""
	}

	for _, user := range usersList {
		userVal, ok := user.(map[string]interface{})
		if !ok {
			continue
		}

		openid, openidok := userVal["openid"].(string)
		if !openidok || openid != wechatopenid {
			continue
		}

		findval, findvalok := userVal[index].(string)
		if findvalok {
			return findval
		}
	}

	return ""
}

// 判断一个 opendid 是否存在
func FindAOpenIDExist(openID string) bool {
	usersList, ok := viper.Get("users").([]interface{})
	if !ok {
		return false
	}

	for _, user := range usersList {
		userVal, ok := user.(map[string]interface{})
		if !ok {
			continue
		}

		opnid, opnidok := userVal["openid"].(string)
		if opnidok && opnid == openID {
			return true
		}
	}

	return false
}
