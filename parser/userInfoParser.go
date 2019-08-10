package parser

import (
	"crawler/engine"
	"regexp"
)

const userInfoRegex = `<div[\s]+class="des f-cl"\s[\w-]+>([\p{Han}]+)\s\|\s([\d]{2}岁)\s\|\s([\p{Han}]+)\s\|\s([\p{Han}]+)\s\|\s([\d]{3}cm)\s\|\s([^<]+)</div>`

func ParseUserInfo(bytes []byte) []engine.Request {
	compile := regexp.MustCompile(userInfoRegex)
	userInfoSubMatches := compile.FindSubmatch(bytes)

	var requests []engine.Request
	requests = append(requests, engine.Request{Name: "地方:" + string(userInfoSubMatches[1]),})
	requests = append(requests, engine.Request{Name: "年龄:" + string(userInfoSubMatches[2]),})
	requests = append(requests, engine.Request{Name: "学历:" + string(userInfoSubMatches[3]),})
	requests = append(requests, engine.Request{Name: "婚姻状况:" + string(userInfoSubMatches[4]),})
	requests = append(requests, engine.Request{Name: "身高:" + string(userInfoSubMatches[5]),})
	requests = append(requests, engine.Request{Name: "收入:" + string(userInfoSubMatches[6]),})
	return requests
}
