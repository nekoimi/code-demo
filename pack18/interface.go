package pack18

import "fmt"

/**
	Golang中的接口无需implements关键字
	只要对象实现了接口的全部方法
	就说这个对象实现了该接口
 */

type Bird interface {
	Fly() string
	Run() string
}

type BirdOne struct {
	Name string
}

func (b BirdOne) Fly() string {
	return b.Name + " flying... "
}

func (b BirdOne) Run() string {
	return b.Name + " running... "
}

func Pack18Main () {
	birdOne := BirdOne{"Bird_One"}
	fmt.Println(birdOne.Fly())
	fmt.Println(birdOne.Run())

	/**
		将对象赋值给接口
		通过接口调用对象的属性方法
	 */
	var bird Bird
	bird = birdOne
	fmt.Println(bird.Fly())
	fmt.Println(bird.Run())
}
