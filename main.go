package main

import (
	"crawler/engine"
	"crawler/parser"
	"log"
)

func main() {
	successfulFetchedCount := engine.Start(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

	log.Printf("successful fetched %d url!", successfulFetchedCount)
}
