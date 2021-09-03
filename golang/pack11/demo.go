package pack11

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
}

type Boss struct {
	Human
	company string
	money float32
}

type Men interface {
	SayHi()
	Sing(msg ...interface{})
}

func (h *Human) SayHi () {
	fmt.Printf("Hi! %s", h.name)
}

func (h *Human) Sing (msg ...interface{}) {
	fmt.Println("la la la la la ....", msg)
}
