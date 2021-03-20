package scheduler

import "goLangLearn/concuImpl/engine"

// SimpleScheduler implements the Scheduler interface defined in engine, So SimpleScheduler is Scheduler
type SimpleScheduler struct {
	workChan chan engine.Request
}

// func (s *SimpleScheduler) ConfigMasterWokerChan(c chan engine.Request) {
// 	s.workChan = c
// }

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workChan <- r
	}()
}
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request)
}

//
