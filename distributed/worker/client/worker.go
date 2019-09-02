package client

import (
	"crawler/distributed/config"
	"crawler/distributed/worker"
	"crawler/engine"
	"net/rpc"
)

// clients []*rpc.Client
func CreateProcessor(clientChan chan *rpc.Client) engine.RequestProcessorFunc {
	return func(req engine.Request) (*engine.ParseResult, error) {
		serializedRequest := worker.SerializeRequest(req)
		var serializedParseResult worker.SerializedParseResult
		c := <-clientChan
		if err := c.Call(config.WorkerRpc, serializedRequest, &serializedParseResult); err == nil {
			return worker.DeserializeParseResult(&serializedParseResult), nil
		} else {
			return nil, err
		}
	}
}
