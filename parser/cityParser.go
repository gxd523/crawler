package parser

import (
	"crawler/distributed/config"
	"crawler/engine"
	"regexp"
)

const (
	// <a href="http://album.zhenai.com/u/1599085703" target="_blank">伟伟</a>
	userListRegex = `<tbody>[^a]+a[\s]+href="(http://album.zhenai.com/u/[\d]+)[^>]+>([^<]+)[^性]+性别：[^\p{Han}]+([\p{Han}])士(.|\n)+?</tbody>`

	// <a href="http://www.zhenai.com/zhenghun/shanghai/2">2</a>
	userListPageRegex = `<a[\s]+href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)">([^<]+)</a>`
)

var (
	userListCompile     = regexp.MustCompile(userListRegex)
	userListPageCompile = regexp.MustCompile(userListPageRegex)
)

func ParseCity(bytes []byte, _ string) *engine.ParseResult {
	userListSubMatches := userListCompile.FindAllSubmatch(bytes, -1)

	parseResult := engine.ParseResult{}
	for _, userSubMatches := range userListSubMatches {
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:  string(userSubMatches[1]),
			Name: string(userSubMatches[2]),
			Parser: &UserInfoParser{
				Username: string(userSubMatches[2]),
				Gender:   string(userSubMatches[3]),
			},
		})
	}

	cityPageSubMatches := userListPageCompile.FindAllSubmatch(bytes, -1)
	for _, citySubMatches := range cityPageSubMatches { // 城市下的用户列表分页
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:    string(citySubMatches[1]),
			Name:   string(citySubMatches[2]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return &parseResult
}
