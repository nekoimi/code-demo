package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main () {
	response, err := http.Get("https://www.dsndsht23.com/forum-103-1.html")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Status Code : ", response.StatusCode)
	}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	re, err := regexp.Compile(``)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(all))
	fmt.Println("===================================")
	res := re.FindAllString(string(all), -1)
	fmt.Println(res)
}
