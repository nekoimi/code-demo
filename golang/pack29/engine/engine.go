package engine

import (
	"../fetcher"
	"fmt"
)

//func Run(seeds ...Request)  {
//	var requests []Request
//	for _, v := range seeds {
//		requests = append(requests, v)
//	}
//
//	for len(requests) > 0 {
//		r := requests[0]
//		requests = requests[1:]
//
//		parseResult, workerErr := worker(r)
//		if workerErr != nil {
//			continue
//		}
//
//		requests = append(requests, parseResult.Requests...)
//		for _, item := range parseResult.Items {
//			log.Printf("DataItem ===> %v", item)
//		}
//	}
//}

func worker(r Request) (ParseResult, error) {
	body, fetchErr := fetcher.Fetch(r.Url)
	if fetchErr != nil {
		fmt.Printf("Fetcher Error.  Url => [ %s ], Error => %v", r.Url, fetchErr)
		return ParseResult{}, fetchErr
	}
	return r.ParseFunc(body), nil
}

