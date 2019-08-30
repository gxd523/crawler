package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcutil"
	"crawler/distributed/worker"
	"crawler/model"
	"fmt"
	"testing"
	"time"
)

func TestWorkerServer(t *testing.T) {
	go func() {
		err := rpcutil.NewServeRpc(fmt.Sprintf(":%d", config.WorkerPort), &worker.CrawlService{})
		if err != nil {
			t.Errorf("error: %v\n", err)
		}
	}()
	time.Sleep(time.Second)

	client, err := rpcutil.NewClient(fmt.Sprintf(":%d", config.WorkerPort))
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	serializedRequest := worker.SerializedRequest{
		Url:  "http://album.zhenai.com/u/1073350475",
		Name: "燕子",
		Parser: worker.SerializedParser{
			FuncName: config.ParseUserInfo,
			Args:     []string{"燕子", "女"},
		},
	}
	var serializedParseResult worker.SerializedParseResult
	err = client.Call(config.WorkerRpc, serializedRequest, &serializedParseResult)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	parseResult := worker.DeserializeParseResult(serializedParseResult)
	userInfo, err := model.FromJsonObj(parseResult.Item.Payload)
	if err != nil {
		t.Errorf("error cast %v\n", parseResult.Item.Payload)
	}
	fmt.Println(userInfo)
}
