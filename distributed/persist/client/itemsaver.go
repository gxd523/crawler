package client

import (
	"context"
	"crawler/distributed/config"
	"crawler/distributed/rpcutil"
	"crawler/engine"
	"errors"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	itemChan := make(chan engine.Item)
	client, err := rpcutil.NewClient(host)
	if err != nil {
		return nil, err
	}
	go func() {
		for item := range itemChan {
			// call RPC to save item
			var result string
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil || result != "ok" {
				log.Printf("ItemSaverRpc Error: %v, result: %s\n", err, result)
				continue
			}
		}
	}()
	return itemChan, nil
}

func Save(item engine.Item, client *elastic.Client, index string) error {
	if item.Id == "" || item.Type == "" {
		return errors.New("id or type is empty")
	}

	_, err := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
