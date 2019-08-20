package persist

import (
	"context"
	"crawler/model"
	"crawler/src/github.com/olivere/elastic"
	"fmt"
	"log"
)

func ItemSaver() chan model.UserInfo {
	itemChan := make(chan model.UserInfo)
	go func() {
		userInfoCount := 0
		for userInfo := range itemChan {
			userInfoCount++
			userInfo.Print(userInfoCount)
			id, err := save(userInfo)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println(id)
		}
	}()
	return itemChan
}

func save(userInfo model.UserInfo) (id string, err error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		return "", err
	}

	response, err := client.Index().
		Index("dating_userinfo").
		Type("zhenai").
		BodyJson(userInfo).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return response.Id, nil
}
