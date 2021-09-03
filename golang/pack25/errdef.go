package main

import "fmt"

/**
	defer  -> 栈  先进后出
 */

func main()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}
