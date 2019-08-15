package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (*SimpleEngine) Start(seeds ...Request) {
	queue := RequestQueue{}
	queue.Push(seeds...)

	count := 0
	for len(queue) > 0 {
		req := queue.Pop()

		log.Printf("left request count=%d, Fetching->%s...%s\n", len(queue), req.Name, req.Url)

		if parseResult, err := work(req); err == nil {
			queue.Push(parseResult.Requests...)

			if parseResult.UserProfile != nil {
				count++
				parseResult.UserProfile.Print(count)
			}
		} else {
			log.Printf("Fetching Error! url=%s error=%v\n", req.Url, err)
			continue
		}
	}
}

func work(req Request) (*ParseResult, error) {
	if bytes, err := fetcher.Fetch(req.Url); err == nil {
		parseResult := req.ParseFunc(bytes)
		return &parseResult, nil
	} else {
		return nil, err
	}
}
