package main

import "fmt"

/**
	Golang 数组
 */

func printArray(arr *[3]int)  {
	arr[0] = 233
	for i, v := range arr{
		fmt.Println(i, v)
	}
}

func Pack24Array () {

	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{2,3,3}
	var arr4[5][5]int
	fmt.Println(arr1, arr2, arr3, arr4)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	printArray(&arr3)

	fmt.Println(arr3)
}
