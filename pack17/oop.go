/**
 * # ------------Oooo---
 * # -----------(----)---
 * # ------------)--/----
 * # ------------(_/-
 * # ----oooO----
 * # ----(---)----
 * # -----\--(--
 * # ------\_)-
 * # ----
 * #     author : Yprisoner <yyprisoner@gmail.com>
 * #     19-5-26 下午9:41
 * #                            ------
 * #    「 涙の雨が頬をたたくたびに美しく 」
 **/
package pack17

import "fmt"

/**
  继承
 */

type Person struct {
	Name string
	Age int8
}

func (p *Person) getName () (name string) {
	return p.Name
}

type Student struct {
	Person
	Class string
	Number int16
}

func (s *Student) setNumber(number int16) {
	s.Number = number
}

func Pack17Main () {
	student := Student{Class: "first"}
	student.setNumber(2333)

	student.Name = "张三"
	student.Age = 20

	fmt.Println("Extends")
	fmt.Println(student.getName())
}
