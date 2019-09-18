package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type ReadyNotifier interface {
	WorkerReady(worker chan Request)
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

func (c *ConcurrentEngine) Run(seeds []Request)  {
	out := make(chan ParseResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerCount ; i++  {
		createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	for _, s := range seeds {
		c.Scheduler.Submit(s)
	}

	for  {
		result := <- out
		for _, item := range result.Items {
			go func() {
				c.ItemChan <- item
			}()
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}

}

func createWorker (in chan Request, out chan ParseResult, reader ReadyNotifier)  {
	go func() {
		for  {
			reader.WorkerReady(in)
			request := <- in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
