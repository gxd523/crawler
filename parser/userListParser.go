package parser

import (
	"crawler/engine"
	"regexp"
)

const userListRegex = `<a[\s]+href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`

func ParseUserList(bytes []byte) engine.ParseResult {
	compile := regexp.MustCompile(userListRegex)
	userListSubMatches := compile.FindAllSubmatch(bytes, -1)

	parseResult := engine.ParseResult{}
	for _, userSubMatches := range userListSubMatches {
		userName := string(userSubMatches[2])
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:  string(userSubMatches[1]),
			Name: userName,
			ParseFunc: func(bytes []byte) engine.ParseResult { // 编程技巧
				return ParseUserInfo(bytes, userName)
			},
		})
	}
	return parseResult
}
