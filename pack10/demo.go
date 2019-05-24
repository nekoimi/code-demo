package pack10

import "fmt"

func Test()  {
	fmt.Println("test")
}


const (
	Sex = iota

	Male
	Female
)

type people struct {
	Name string
	Age int
	Sex int
}

func (p people) SayHello()  {
	fmt.Printf("I am %s , %d years old this year !\n", p.Name, p.Age)
}

func SayHelloT(p people) {
	fmt.Printf("I am %s , %d years old this year !\n", p.Name, p.Age)
}

func mainTT()  {
	people1 := people{
		"张三",
		18,
		Male,
	}

	people2 := people{
		"李四",
		20,
		Female,
	}

	people3 := people{
		"王二麻子",
		9,
		Female,
	}


	people1.SayHello()
	people2.SayHello()
	people3.SayHello()
	SayHelloT(people3)

}

