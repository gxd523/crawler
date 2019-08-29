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
		Id:   "1073350475",
		Url:  "http://album.zhenai.com/u/1073350475",
		Payload: model.UserInfo{
			Name:        "燕子",
			Gender:      "女",
			Age:         33,
			Height:      166,
			Weight:      59,
			Birthplace:  "甘肃平凉",
			Xinzuo:      "摩羯座",
			City:        "伊犁哈萨克自治州伊宁",
			Nationality: "回族",
			Education:   "高中及以下",
			Marriage:    "离异",
			Occupation:  "",
			Income:      "5-8千",
			House:       "已购房",
			Car:         "未买车",
			HaveChild:   "有孩子且住在一起",
			WannaChild:  "视情况而定",
		},
	}

	const index = "dating_test"

	// TODO Try to start up elastic search(here using docker go client.)
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = Save(expectedItem, client, index)
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
