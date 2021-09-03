package main

import "fmt"

func showSliceInfo(s []int)  {
	fmt.Printf("slice => %v\t length=%d\t cap=%d \n",s, len(s), cap(s))
}

func Demo()  {
	arr := [...]int{0,1,2,3,4,5,6,7,8,9}
	s1 := arr[2:6]  // 2,3,4,5
	s2 := s1[3:5]   // 5,6
	fmt.Println(s1)
	/**

	 */
	fmt.Println(s2)

	showSliceInfo(s1)
	showSliceInfo(s2)
	s3 := append(s1, 233)
	showSliceInfo(s3)
}

func Pack24Slice()  {
	//arr := [...]int{1,2,3,4,5,6,7}
	//fmt.Println(arr[2:5])
	Demo()
}

