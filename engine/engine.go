package engine

import (
	"crawler/fetcher"
	"crawler/model"
	"log"
)

func Start(seeds ...Request) int {
	queue := RequestQueue{}
	queue.Push(seeds...)

	var successfulFetchedCount int

	for len(queue) > 0 {
		req := queue.Pop()

		log.Printf("left request count=%d, Fetching->%s...%s\n", len(queue), req.Name, req.Url)
		bytes, err := fetcher.Fetch(req.Url)
		if err != nil {
			log.Printf("Fetching Error! url=%s error=%v\n", req.Url, err)
			continue
		}
		successfulFetchedCount++
		parseResult := req.ParseFunc(bytes)
		queue.Push(parseResult.Requests...)

		emptyUserInfo := model.UserInfo{}
		if parseResult.UserProfile != emptyUserInfo {
			parseResult.UserProfile.Print()
		}
	}

	return successfulFetchedCount
}
