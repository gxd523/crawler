package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcutil"
	"crawler/distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatalln(rpcutil.NewServeRpc(fmt.Sprintf(":%d", config.WorkerPort), &worker.CrawlService{}))
}
