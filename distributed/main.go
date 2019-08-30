package main

import (
	"crawler/distributed/config"
	itemSaverClient "crawler/distributed/persist/client"
	workerClient "crawler/distributed/worker/client"
	"crawler/engine"
	"crawler/parser"
	"crawler/scheduler"
	"fmt"
)

func main() {
	itemChan, err := itemSaverClient.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	processor, err := workerClient.CreateProcessor()
	if err != nil {
		panic(err)
	}
	myEngine := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      99,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	myEngine.Start(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Name:   "城市列表",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList)})
}
