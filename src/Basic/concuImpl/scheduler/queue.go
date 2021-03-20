package scheduler

import "goLangLearn/concuImpl/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	// engine 中的每个 worker 都 make(chan Request)，因此 worker 是 chan Request 类型
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		// 创建 Request 和 Worker 的队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			// 每次循环处理三种情况中的一种，不然就卡住
			// activeRequest 和 activeWorker 每次循环创建
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
