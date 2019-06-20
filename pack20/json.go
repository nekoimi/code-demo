package pack20

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	UserName string 	`json:"user_name"`
	Age int8			`json:"age"`
}

func Pack20() {
	arr := [5]int{1,4,7,2,5}
	byte1, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte1))

	map1 := make(map[string]string)
	map1["username"] = "zhangsan"
	map1["aaa"] = "xxxxxx"
	byte2, err := json.Marshal(map1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte2))

	ui := UserInfo{"zhangsan", 20}
	byte3, err := json.Marshal(ui)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte3))
}
