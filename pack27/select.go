package main

import (
	"fmt"
)

/**
多个 channel 调度
*/

func generator () chan int {
	o := make(chan int)
	go func() {
		i := 0
		for  {
			o <- i
			i ++
		}
	}()
	return o
}

func main() {
	var c1, c2 = generator(), generator()
	for {
		select {
			case n := <-c1:
				fmt.Println("form c1", n)
			case n := <-c2:
				fmt.Println("form c2", n)
		}
	}
}
