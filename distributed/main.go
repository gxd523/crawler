package main

import (
	"crawler/distributed/config"
	itemSaverClient "crawler/distributed/persist/client"
	"crawler/distributed/rpcutil"
	workerClient "crawler/distributed/worker/client"
	"crawler/engine"
	"crawler/parser"
	"crawler/scheduler"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemSaverClient.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	clientChan := createClientPool(strings.Split(*workerHosts, ","))
	processor := workerClient.CreateProcessor(clientChan)
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

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		if client, err := rpcutil.NewClient(host); err == nil {
			clients = append(clients, client)
		} else {
			log.Printf("error connecting to %s: %v", host, err)
		}
	}

	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				clientChan <- client
			}
		}
	}()
	return clientChan
}
