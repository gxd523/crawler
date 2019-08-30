package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
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
		if parseResult.Item != nil {
			go func() { engine.ItemChan <- *parseResult.Item }()
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
			if parseResult, err := Work(req); err == nil {
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
