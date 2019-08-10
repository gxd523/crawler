package main

import (
	"demos/crawler/engine"
	"demos/crawler/parser"
	"log"
)

func main() {
	successfulFetchedCount := engine.Start(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

	log.Printf("successful fetched %d url!", successfulFetchedCount)

	//engine.Start(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/aba",
	//	Name:      "阿坝",
	//	ParseFunc: parser.ParseUserList,
	//})
}
