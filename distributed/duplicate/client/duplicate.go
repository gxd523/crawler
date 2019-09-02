package client

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcutil"
	"crawler/engine"
	"log"
)

func CreateIsDuplicateUrlFunc(host string) engine.IsDuplicateUrlFunc {
	client, err := rpcutil.NewClient(host)
	if err != nil {
		log.Printf("error: %v", err)
	}
	return func(url string) bool {
		result := true
		err := client.Call(config.DuplicateRpc, url, &result)
		if err != nil {
			return false
		}
		return result
	}
}
