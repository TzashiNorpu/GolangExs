package engine

import (
	"goLangLearn/concuImpl/fetcher"
	"log"
)

type SimpleEngine struct{}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		//log.Printf("Fetching %s", r.Url)
		//body, err := fetcher.Fetch(r.Url)
		//if err != nil {
		//	log.Printf("Fetcher: error "+"fetching url%s : %v", r.Url, err)
		//	continue
		//}
		//parseResult := r.ParserFunc(body)
		//--抽成worker
		parseResult, _ := simplework(r)

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}

func simplework(r Request) (ParseResult, error) {
	// log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url%s : %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
