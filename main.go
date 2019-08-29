package main

import (
	"crawler/config"
	"crawler/engine"
	"crawler/parser"
	"crawler/persist"
	"crawler/scheduler"
	"fmt"
)

func main() {
	itemChan, err := persist.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	myEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 99,
		ItemChan:    itemChan,
	}

	myEngine.Start(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		Name:      "城市列表",
		ParseFunc: parser.ParseCityList})
}