package parser

import (
	"demos/crawler/engine"
	"regexp"
)

// <a href="http://www.zhenai.com/zhenghun/sanya" data-v-5e16505f>三亚</a>
const cityListRegex = `<a[\s]+href="(http://www.zhenai.com/zhenghun/[\w]+)" [^>]*>([^<]+)</a>`

func ParseCityList(bytes []byte) []engine.Request {
	compile := regexp.MustCompile(cityListRegex)
	cityListSubMatches := compile.FindAllSubmatch(bytes, -1)

	var requests []engine.Request
	for _, citySubMatches := range cityListSubMatches {
		requests = append(requests, engine.Request{
			Url:       string(citySubMatches[1]),
			Name:      string(citySubMatches[2]),
			ParseFunc: ParseUserList,
		})
	}

	return requests
}
