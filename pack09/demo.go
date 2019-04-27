package pack09

import (
	"fmt"
	"time"
)

/**
	消费者
 */
func consumer(data chan int, done chan bool)  {
	for x := range data  {
		fmt.Printf("consumer %d\n", x)
	}

	done <- true
}

/**
	生产者
 */
func producer(data chan int)  {
	index := 0
	for index < 100 {
		data <- index
		time.Sleep(time.Second)
		index ++
	}

	close(data)
}

func TestChannel()  {
	done := make(chan bool)  // 用于接收消费结束
	data := make(chan int)   // 数据管道

	go consumer(data, done)
	go producer(data)

	<-done  // 阻塞
}
