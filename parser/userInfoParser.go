package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

const (
	// <div class="des f-cl" data-v-3c42fade>阿坝 | 26岁 | 大学本科 | 离异 | 156cm | 12001-20000元</div>
	userInfoRegex = `<div[\s]+class="des f-cl"\s[\w-]+>([\p{Han}]+)\s\|\s([\d]{2})岁\s\|\s([\p{Han}]+)\s\|\s([\p{Han}]+)\s\|\s([\d]{3})cm\s\|\s([^<]+)</div>`

	// http://album.zhenai.com/u/1599085703
	idRegex = `http://album.zhenai.com/u/([\d]+)`
)

var (
	userInfoCompile = regexp.MustCompile(userInfoRegex)
	idCompile       = regexp.MustCompile(idRegex)
)

func ParseUserInfo(bytes []byte, name string, url string) engine.ParseResult {
	userInfoSubMatches := userInfoCompile.FindSubmatch(bytes)

	userInfo := model.UserInfo{}
	userInfo.Name = name
	if userInfoSubMatches != nil {
		if age, err := strconv.Atoi(string(userInfoSubMatches[2])); err == nil {
			userInfo.Age = age
		}
		userInfo.Occupation = string(userInfoSubMatches[1])
		userInfo.Education = string(userInfoSubMatches[3])
		userInfo.Marriage = string(userInfoSubMatches[4])
		if height, err := strconv.Atoi(string(userInfoSubMatches[5])); err == nil {
			userInfo.Height = height
		}
		userInfo.Income = string(userInfoSubMatches[6])
	}
	item := engine.Item{
		Type:    "zhenai",
		Id:      idCompile.FindStringSubmatch(url)[1],
		Url:     url,
		Payload: userInfo,
	}
	return engine.ParseResult{Item: &item}
}
