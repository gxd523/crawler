package main

import (
	"crawler/distributed/rpcutil"
	"crawler/persist"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func main() {
	// 相比panic()，fatal()连recover()的机会都不给，直接退出程序
	log.Fatalln(newItemSaverServeRpc(":1234", "dating_userinfo"))
}

func newItemSaverServeRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcutil.NewServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
