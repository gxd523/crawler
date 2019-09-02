package main

import (
	"crawler/distributed/rpcutil"
	"crawler/distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for me listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatalln(rpcutil.NewServeRpc(fmt.Sprintf(":%d", *port), &worker.CrawlService{}))
}
