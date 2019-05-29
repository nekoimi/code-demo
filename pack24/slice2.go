package main

import "fmt"

func Pack24Slice2()  {
	var s []int
	fmt.Println(s)
	for i := 0; i < 100 ;i++  {
		s = append(s, i * 2 + 1)
	}
	fmt.Println(s)

	s1 := []int{2,3,3}
	showSliceInfo(s1)

	s2 := make([]int, 10)
	showSliceInfo(s2)

	s3 := make([]int, 10, 16)
	showSliceInfo(s3)

	copy(s2, s1)
	showSliceInfo(s2)

	s2 = append(s2[:1], s1[:]...)
	showSliceInfo(s2)
}
