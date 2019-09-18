package main

import "fmt"

var (
	var_int = 3
	var_str = "233"
)


func defaultVar () {
	var a int
	var b string
	fmt.Printf("%d %q", a, b)
	var c, d, e = 233, "233", true
	println(c, d, e)
}

const (
	const_var1 = 233
)

func enums()  {

	const (
		php = 0
		java = 2
		c = 3
		golang = 4
	)

	println(php, java, c, golang)
}

func main() {
	println(var_int)
	println(var_str)
	defaultVar()
	println(const_var1)
}
