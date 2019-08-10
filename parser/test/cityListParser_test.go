package test

import (
	"demos/crawler/parser"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	bytes, _ := ioutil.ReadFile("cityListParser_data.html")
	parseResult := parser.ParseCityList(bytes)

	const expectedCount = 470
	if len(parseResult.Requests) != expectedCount {
		t.Errorf("Test Failed! expect have %d Request, actual have %d\n", expectedCount, len(parseResult.Requests))
	}
	if len(parseResult.Names) != expectedCount {
		t.Errorf("Test Failed! expect have %d Item, actual have %d\n", expectedCount, len(parseResult.Names))
	}

	expectUrlMap := map[int]string{
		52:  "http://www.zhenai.com/zhenghun/chengdu",
		313: "http://www.zhenai.com/zhenghun/shaoguan",
		464: "http://www.zhenai.com/zhenghun/zhumadian",
	}
	expectItemMap := map[int]string{
		52:  "成都",
		313: "韶关",
		464: "驻马店",
	}

	for k, url := range expectUrlMap {
		if parseResult.Requests[k].Url != url {
			t.Errorf("Test Failed! expect %d url is %s, actual is %s\n", k, url, parseResult.Requests[k].Url)
		}
	}

	for k, item := range expectItemMap {
		if parseResult.Names[k] != item {
			t.Errorf("Test Failed! expect %d item is %s, actual is %s\n", k, item, parseResult.Names[k])
		}
	}
}
