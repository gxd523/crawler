package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"gopkg.in/olivere/elastic.v6"
	"testing"
)

func TestSave(t *testing.T) {
	expectedItem := engine.Item{
		Type: "zhenai",
		Id:   "",
		Url:  "",
		Payload: model.UserInfo{
			Name:       "郭晓东",
			Gender:     "male",
			Age:        28,
			Height:     176,
			Income:     "1-2W",
			Marriage:   "married",
			Education:  "college",
			Occupation: "Hangzhou",
			Birthplace: "Xinjiang",
			House:      "Yes",
			Car:        "No",
			Xinzuo:     "Gemini",
		},
	}

	const index = "dating_test"

	// TODO Try to start up elastic search(here using docker go client.)
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = save(expectedItem, client, index)
	if err != nil {
		panic(err)
	}

	result, err := client.Get().
		Index(expectedItem.Id).
		Type(expectedItem.Type).
		Id(expectedItem.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actualItem engine.Item
	err = json.Unmarshal(*result.Source, &actualItem)
	if err != nil {
		panic(err)
	}

	// 别扭操作
	actualUserInfo, _ := model.FromJsonObj(actualItem.Payload)
	actualItem.Payload = actualUserInfo

	if expectedItem != actualItem {
		t.Errorf("got %+v, expected %v", actualItem, expectedItem)
	}
}
