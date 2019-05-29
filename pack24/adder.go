package main

import (
	"fmt"
	"strconv"
)

/**
	函数式编程
	- 不可变性
	- 函数只有一个参数
 */

func demo () func(int) string  {
	return func(i int) string {
		return "i is " + strconv.Itoa(i)
	}
}

func main()  {
	d := demo()
	fmt.Println(d(233))
}
