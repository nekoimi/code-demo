package main

import "fmt"

/**
	Golang Map -> 无序！
 */

func Pack24Map()  {
	m1 := map[string]string {
		"username": "demo100",
		"password": "demo100",
		"demo001" : "demo001",
		"demo002" : "demo002",
		"demo003" : "demo003",
		"demo004" : "demo004",
		"demo005" : "demo005",
		"demo006" : "demo006",
	}

	m2 := make(map[string]string)

	fmt.Println(m1, m2)

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	username, ok := m1["username"]
	fmt.Println(username, ok)
	if not_exists, ok := m1["not_exists"]; !ok {
		fmt.Println("不存在")
	} else {
		fmt.Println(not_exists, ok)
	}

	delete(m1, "username")  // 删除
	fmt.Println(m1)
}
