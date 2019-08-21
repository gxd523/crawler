package engine

import (
	"crawler/model"
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

			if parseResult.Item != nil {
				count++
				userInfo := parseResult.Item.Payload.(model.UserInfo)
				userInfo.Print(req.Url, count)
			}
		} else {
			log.Printf("Fetching Error! url=%s error=%v\n", req.Url, err)
			continue
		}
	}
}
