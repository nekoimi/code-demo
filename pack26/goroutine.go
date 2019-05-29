package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	var arr [10]int
	for i := 0;i < 10 ;i++  {
		go func(i int) {
			for  {
				arr[i] ++
				runtime.Gosched()
				//fmt.Println("hello ", i)
			}
		}(i)
	}

	time.Sleep(time.Microsecond)
	fmt.Println(arr)
}
