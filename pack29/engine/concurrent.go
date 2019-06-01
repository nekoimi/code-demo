package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChan (chan Request)
}

var Count = 0

func (c *ConcurrentEngine) Run(seeds []Request)  {
	in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < c.WorkerCount ; i++  {
		createWorker(in ,out)
	}

	for _, s := range seeds {
		c.Scheduler.Submit(s)
	}

	for  {
		result := <- out
		log.Printf("DataItem Count ===> %v", Count)
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
		Count ++
	}

}

func createWorker (in chan Request, out chan ParseResult)  {
	go func() {
		for  {
			request := <- in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
