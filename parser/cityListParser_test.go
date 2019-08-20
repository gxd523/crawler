package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	bytes, _ := ioutil.ReadFile("cityListParser_data.html")
	parseResult := ParseCityList(bytes)

	const expectedCount = 470
	if len(parseResult.Requests) != expectedCount {
		t.Errorf("Test Failed! expect have %d Request, actual have %d\n", expectedCount, len(parseResult.Requests))
	}

	expectUrlMap := map[int]string{
		52:  "http://www.zhenai.com/zhenghun/chengdu",
		313: "http://www.zhenai.com/zhenghun/shaoguan",
		464: "http://www.zhenai.com/zhenghun/zhumadian",
	}

	for k, url := range expectUrlMap {
		if parseResult.Requests[k].Url != url {
			t.Errorf("Test Failed! expect %d url is %s, actual is %s\n", k, url, parseResult.Requests[k].Url)
		}
	}
}
