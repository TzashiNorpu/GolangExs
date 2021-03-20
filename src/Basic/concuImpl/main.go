package main

import (
	"goLangLearn/concuImpl/engine"
	"goLangLearn/concuImpl/persist"
	"goLangLearn/concuImpl/scheduler"
	"goLangLearn/concuImpl/zhenai/parser"
)

func main() {
	// concurrent with scheduler
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	// url := "http://localhost:8080/mock/www.zhenai.com/zhenghun"
	// e.Run(engine.Request{
	// 	Url:        url,
	// 	ParserFunc: parser.ParseCityList,
	// })

	url := "http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai"
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCity,
	})
	// simple engine without scheduler
	// s := engine.SimpleEngine{}
	// s.Run(engine.Request{
	// 	Url:        url,
	// 	ParserFunc: parser.ParseCityList,
	// })
}
