package client

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcutil"
	"crawler/distributed/worker"
	"crawler/engine"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	if client, err := rpcutil.NewClient(fmt.Sprintf(":%d", config.WorkerPort)); err == nil {
		return func(req engine.Request) (*engine.ParseResult, error) {
			serializedRequest := worker.SerializeRequest(req)
			var serializedParseResult worker.SerializedParseResult
			if err := client.Call(config.WorkerRpc, serializedRequest, &serializedParseResult); err == nil {
				return worker.DeserializeParseResult(&serializedParseResult), nil
			} else {
				return nil, err
			}
		}, nil
	} else {
		return nil, err
	}
}
