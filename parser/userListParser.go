package parser

import (
	"crawler/engine"
	"regexp"
)

const (
	// <a href="http://album.zhenai.com/u/1599085703" target="_blank">伟伟</a>
	userListRegex = `<a[\s]+href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`

	// <a href="http://www.zhenai.com/zhenghun/shanghai/2">2</a>
	userListPageRegex = `<a[\s]+href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)">([^<]+)</a>`
)

var (
	userListCompile     = regexp.MustCompile(userListRegex)
	userListPageCompile = regexp.MustCompile(userListPageRegex)
)

func ParseUserList(bytes []byte) engine.ParseResult {
	userListSubMatches := userListCompile.FindAllSubmatch(bytes, -1)

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

	cityPageSubMatches := userListPageCompile.FindAllSubmatch(bytes, -1)
	for _, citySubMatches := range cityPageSubMatches { // 城市下的用户列表分页
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:       string(citySubMatches[1]),
			Name:      string(citySubMatches[2]),
			ParseFunc: ParseUserList,
		})
	}
	return parseResult
}
