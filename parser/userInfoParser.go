package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

const userInfoRegex = `<div[\s]+class="des f-cl"\s[\w-]+>([\p{Han}]+)\s\|\s([\d]{2})岁\s\|\s([\p{Han}]+)\s\|\s([\p{Han}]+)\s\|\s([\d]{3})cm\s\|\s([^<]+)</div>`

var userInfoCompile = regexp.MustCompile(userInfoRegex)

func ParseUserInfo(bytes []byte, name string) engine.ParseResult {
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
	return engine.ParseResult{UserProfile: userInfo,}
}
