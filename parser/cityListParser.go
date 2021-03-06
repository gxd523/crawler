package parser

import (
	"crawler/distributed/config"
	"crawler/engine"
	"regexp"
)

// <a href="http://www.zhenai.com/zhenghun/sanya" data-v-5e16505f>三亚</a>
const cityListRegex = `<a[\s]+href="(http://www.zhenai.com/zhenghun/[\w]+)" [^>]*>([^<]+)</a>`

func ParseCityList(bytes []byte, _ string) *engine.ParseResult {
	compile := regexp.MustCompile(cityListRegex)
	cityListSubMatches := compile.FindAllSubmatch(bytes, -1)

	parseResult := engine.ParseResult{}
	for _, citySubMatches := range cityListSubMatches {
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:    string(citySubMatches[1]),
			Name:   string(citySubMatches[2]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return &parseResult
}
