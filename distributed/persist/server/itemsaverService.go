package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/rpcutil"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

var port = flag.Int("port", 0, "the port for me listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	// 相比panic()，fatal()连recover()的机会都不给，直接退出程序
	log.Fatalln(newItemSaverServeRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func newItemSaverServeRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcutil.NewServeRpc(host, &persist.ItemSaverService{ // 注意这里要取地址，因为Receiver也是指针类型
		Client: client,
		Index:  index,
	})
}
