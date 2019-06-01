package engine

import (
	"../fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _, v := range seeds {
		requests = append(requests, v)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, fetchErr := fetcher.Fetch(r.Url)
		if fetchErr != nil {
			fmt.Printf("Fetcher Error.  Url => [ %s ], Error => %v", r.Url, fetchErr)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("GoodsName ===> %v", item)
		}
	}
}
