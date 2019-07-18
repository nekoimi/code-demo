package main

import "fmt"

var (
	Ga int = 99
)

const (
	v int = 199
)

func GetGa() func() int {

	/**
		这里定义的是局部变量作用域名是这个外面的 {}
	 */
	{
		if Ga := 55; Ga < 60 {
			fmt.Println("GetGa if 中：", Ga)
		}
	}


	/**
		这里定义的是局部变量作用域名是这个外面的 {}
	*/
	{
		for Ga := 2; ; {
			fmt.Println("GetGa循环中：", Ga)
			break
		}
	}

	/**
		这是使用全局变量
	*/
	fmt.Println("GetGa函数中：", Ga)

	return func() int {
		Ga += 1
		return Ga
	}
}

func main() {
	Ga := "string"
	fmt.Println("main函数中：", Ga)

	b := GetGa()
	fmt.Println("main函数中：", b(), b(), b(), b())

	v := 1
	{
		v := 2
		fmt.Println(v)
		{
			v := 3
			fmt.Println(v)
		}
	}
	fmt.Println(v)
}
