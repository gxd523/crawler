package main

import (
	"crawler/config"
	"crawler/distributed/rpcutil"
	"crawler/engine"
	"crawler/model"
	"fmt"
	"testing"
	"time"
)

func TestItemSaverServeRpc(t *testing.T) {
	go newItemSaverServeRpc(fmt.Sprintf(":%d", config.ItemSaverPort), "test1")

	// TODO
	time.Sleep(time.Second)

	client, _ := rpcutil.NewClient(fmt.Sprintf(":%d", config.ItemSaverPort))

	item := engine.Item{
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
	var result string

	err := client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("expected return ok, actual return %s\n", result)
	}
}
