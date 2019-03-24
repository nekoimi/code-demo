package main

import "./pack03"

func main(){

	//var nums = []int {1,8,9,5,6,1,5,2,6,4,7,5,5}
	//
	//var result []int = pack01.TwoSum1(nums, 10)
	//for i := 0; i< len(result) ; i++  {
	//	println(result[i])
	//}

	var arr = [][] int {
		{1, 3, 6, 11},
		{2, 7, 9, 13},
		{5, 10, 12, 20},
	}

	result := pack03.ArrayFind(arr, 10)
	println(result)
}
