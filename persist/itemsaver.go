package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"errors"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	itemChan := make(chan engine.Item)
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}
	go func() {
		userInfoCount := 0
		for item := range itemChan {
			userInfoCount++
			userInfo := item.Payload.(model.UserInfo)
			userInfo.Print(item.Url, userInfoCount)
			err := Save(item, client, index)
			if err != nil {
				log.Println(err)
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
