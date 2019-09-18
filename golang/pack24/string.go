package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	str := "Golang大法好!"
	for _, b := range []byte(str) {
		fmt.Printf("%X ", b)
	}
	fmt.Printf("\n%X\n", str)

	for i, s := range str {
		fmt.Printf("(%d %X) ",i, s)
	}
	fmt.Println("")
	fmt.Println(
		utf8.RuneCountInString(str))

	fmt.Println("")
	for i, s := range []rune(str) {
		fmt.Printf("(%d %c) ",i, s)
	}
	fmt.Println("")
}
