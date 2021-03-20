package engine

import (
	"Basic/concuImpl/fetcher"
	"log"
)

// urls parser hierarchy
// code parserA to parser urlA get urlBs
// 	define urlB get logic and parserB name in parserA
// code parserB to parser urlBs get urlCs
//  define urlC get logic and parserC name in parserB
// assign a url to urlA as entry
// regular expression implements urlBs,urlCs... get logic

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

// type SimpleScheduler struct{...} implement Scheduler interface -> Scheduler 类型
type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	// ConfigMasterWokerChan(chan Request)
	ReadyNotifier
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		// engine 不知道每个 worker 的 chan 情况，因此找 Scheduler 要
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}
	// profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			// if _, ok := item.(model.Profile); ok {
			// 	log.Printf("Got profile:%d  %v", profileCount, item)
			// 	profileCount++
			// }
			go func() {
				e.ItemChan <- item
			}()

		}

		for _, request := range result.Requests {
			// URL dedup
			if isDuplicate(request.Url) {
				// log.Printf("Duplicate request: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	// chan per worker
	// in := make(chan Request)
	go func() {
		for {
			// worker ready
			ready.WorkerReady(in)
			request := <-in
			result, _ := worker(request)
			out <- result
		}
	}()
}

func worker(r Request) (ParseResult, error) {
	// log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s : %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
