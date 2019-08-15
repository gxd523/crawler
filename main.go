package main

import (
	"crawler/engine"
	"crawler/parser"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	myEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 99,
		ItemChan:    persist.ItemSaver(),
	}

	myEngine.Start(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/shanghai",
		Name:      "上海",
		ParseFunc: parser.ParseUserList,})
}