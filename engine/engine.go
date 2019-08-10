package engine

import (
	"crawler/fetcher"
	"fmt"
	"log"
)

func Start(seeds ...Request) int {
	queue := RequestQueue{}
	queue.Push(seeds...)

	var successfulFetchedCount int

	for len(queue) > 0 {
		req := queue.Pop()
		if req.Url == "" {
			continue
		}
		log.Printf("left request count=%d, Fetching->%s...%s\n", len(queue), req.Name, req.Url)
		bytes, err := fetcher.Fetch(req.Url)
		if err != nil {
			log.Printf("Fetching Error! url=%s error=%v\n", req.Url, err)
			continue
		}
		successfulFetchedCount++
		requests := req.ParseFunc(bytes)
		queue.Push(requests...)

		fmt.Print("[")
		for _, req := range requests {
			fmt.Printf("%s ", req.Name)
		}
		fmt.Println("]")
	}

	return successfulFetchedCount
}
