package main

import "fmt"

func lengthNoRepact(s string) int {
	lastO := make(map[byte]int)
	start := 0
	maxL := 0
	for i, char := range []byte(s) {
		if chI, ok := lastO[char]; ok && chI >= start {
			start = chI + 1
		}
		if i - start + 1 > maxL {
			maxL = i - start + 1
		}
		lastO[char] = i
	}
	return maxL
}

func PackMap2()  {
	s1 := "fsddasdasdsdwwadaw"
	s2 := "abcab"
	s3 := "aaa"

	fmt.Println(lengthNoRepact(s1))
	fmt.Println(lengthNoRepact(s2))
	fmt.Println(lengthNoRepact(s3))
}