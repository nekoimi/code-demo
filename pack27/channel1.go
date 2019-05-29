package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	done func()
}

func doWorker(id int, w worker)  {
	for n := range w.in {
		fmt.Printf("id=%d, value=%c\n", id, n)
		w.done()
	}
}

func createWorker(i int, done func()) worker{
	w := worker{
		make(chan int),
		done,
	}
	go doWorker(i, w)
	return w
}

func channelDemo()  {
	// var ch chan int  // init ch == nil
	var wg sync.WaitGroup
	wg.Add(30)

	var ws [10]worker
	for i := 0; i < 10; i++{
		ws[i] = createWorker(i, func() {
			wg.Done()
		})
	}

	for i := 0; i < 10 ;i ++  {
		ws[i].in <- 'a' + i
	}

	for i := 0; i < 10 ;i ++  {
		ws[i].in <- 'A' + i
	}

	for i := 0; i < 10 ;i ++  {
		ws[i].in <- 'a' + i * 2 + 1
	}

	wg.Wait()
	//for _, w := range ws {
	//	<- w.done
	//	<- w.done
	//}
}

func main()  {
	channelDemo()
}
