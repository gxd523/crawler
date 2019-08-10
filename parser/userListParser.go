package parser

import (
	"demos/crawler/engine"
	"regexp"
)

const userListRegex = `<a[\s]+href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`

func ParseUserList(bytes []byte) []engine.Request {
	compile := regexp.MustCompile(userListRegex)
	userListSubMatches := compile.FindAllSubmatch(bytes, -1)

	var requests []engine.Request
	for _, userSubMatches := range userListSubMatches {
		requests = append(requests, engine.Request{
			Url:       string(userSubMatches[1]),
			Name:      string(userSubMatches[2]),
			ParseFunc: ParseUserInfo,
		})
	}
	return requests
}
