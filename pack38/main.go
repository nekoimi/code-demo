package main

import (
	"../../gopkg.in/ini.v1"
	"fmt"
)

func main() {
	config, err := ini.Load("my.ini")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(config.Section("").Key("app_mode").String())
	fmt.Println(config.Section("http").Key("http_host").String())
	keys := config.Section("mysql").Keys()
	for _, key := range keys {
		fmt.Println(key)
	}
}
