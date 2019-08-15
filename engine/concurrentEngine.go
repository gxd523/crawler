package engine

import (
	"crawler/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan model.UserInfo
}

var visitedUrlMap = make(map[string]bool)

func (engine *ConcurrentEngine) Start(seeds ...Request) {
	resultChan := make(chan ParseResult)
	engine.Scheduler.Run()

	for i := 0; i < engine.WorkerCount; i++ { // 开WorkerCount个协程
		doWork(engine.Scheduler, engine.Scheduler.WorkerChan(), resultChan)
	}

	engine.submitRequest(seeds)

	for parseResult := range resultChan {
		if parseResult.UserProfile != nil {
			go func() { engine.ItemChan <- *parseResult.UserProfile }()
		}

		engine.submitRequest(parseResult.Requests)
	}
}

func (engine *ConcurrentEngine) submitRequest(requests []Request) {
	for _, req := range requests {
		if isDuplicateUrl(req.Url) {
			//log.Printf("Duplicate request...%s", req.Url)
		} else {
			engine.Scheduler.Submit(req)
		}
	}
}

func isDuplicateUrl(url string) bool { // TODO md5
	if isVisited := visitedUrlMap[url]; isVisited {
		return true
	} else {
		visitedUrlMap[url] = true
		return false
	}
}

func doWork(notifier WorkerReadyNotifier, requestChan chan Request, resultChan chan ParseResult) {
	go func() {
		for {
			notifier.WorkerReady(requestChan)
			req := <-requestChan
			log.Printf("Fetching->%s...%s\n", req.Name, req.Url)
			if parseResult, err := work(req); err == nil {
				resultChan <- *parseResult
			} else {
				log.Print(err)
				continue
			}
		}
	}()
}

type Scheduler interface { // 用的地方写接口
	Submit(Request)

	WorkerChan() chan Request

	WorkerReadyNotifier

	Run()
}

type WorkerReadyNotifier interface {
	WorkerReady(w chan Request)
}
