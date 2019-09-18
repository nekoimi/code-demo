package pack07

import "fmt"

func TestSlice() {
	var sliceX = make([]int,0,10)
	// sliceX := make([]int,0,10)
	for i := 0;i < 10 ;i++  {
		sliceX = append(sliceX, i)
	}
	fmt.Println(sliceX)
}

func TestMap()  {
	var mapX = make(map[string]string)
	// mapX := make(map[string]string)
	mapX["php"] = "世界上最好的语言."
	mapX["golang"] = "正在学习的语言"

	/**
		ok-idiom模式
		通过返回ok的bool值确定返回成功
	 */
	if value, ok := mapX["php"]; ok {
		fmt.Println(value)
	}

	if value, ok := mapX["not_found"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found.")
	}
}


type UserInfo struct {
	Name string
	Age int
	AvatarUrl string
	XXX bool
}

type User struct {
	UserInfo
	Username string
	Password string
}

func TestUser()  {
	var user User

	user.Username = "demo100"
	user.Password = "demo100"
	user.Name = "Yprisoner"

	fmt.Println(user)
}

type MyInt int

func (myint *MyInt) TestFunc()  {
	*myint++
}


func (userinfo UserInfo) ToString() string {
	return fmt.Sprintf("%v", userinfo)
}

/**
	类型实现接口
 */
type Printer interface {
	Print() string
}

func (user User) Print() string {
	return fmt.Sprintf("%v", user)
}

/***
	空接口类型 -> 相当于Object -> 可以接受任意类型的 TODO 对象
 */
type Object interface {}
