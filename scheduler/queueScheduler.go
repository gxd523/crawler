package scheduler

import "crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request      // 用于requestQueue队列
	workerChan  chan chan engine.Request // 用于workerQueue队列,chan engine.Request用于发送request给doWork
}

func (q *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Submit(req engine.Request) {
	q.requestChan <- req
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueueScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}

			select {
			case r := <-q.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-q.workerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
