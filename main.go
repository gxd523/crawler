package main

import (
	"crawler/distributed/config"
	"crawler/engine"
	"crawler/parser"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_userinfo")
	if err != nil {
		panic(err)
	}
	myEngine := engine.ConcurrentEngine{
		Scheduler:            &scheduler.QueueScheduler{},
		WorkerCount:          99,
		ItemChan:             itemChan,
		RequestProcessorFunc: engine.Work,
		IsDuplicateUrlFunc:   engine.IsDuplicateUrl,
	}

	myEngine.Start(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Name:   "城市列表",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList)})
}
