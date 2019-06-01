package main

import (
	"../engine"
	"../meilele/parser"
	"../scheduler"
	"strconv"
)

func main()  {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	var requests []engine.Request
	for i := 1; i < 47 ; i++  {
		requests = append(requests, engine.Request{
			Url: "https://www.meilele.com/category-keting/list-p"+strconv.Itoa(i)+"/?from=page",
			ParseFunc: parser.ParseGoodList,
		})
	}

	e.Run(requests)
}