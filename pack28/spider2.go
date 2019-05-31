package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main()  {
	response, err := http.Get("http://www.zhenai.com/zhenghun/wuhan")
	if err != nil {
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("status code ===> ", response.StatusCode)
		return
	}
	defer response.Body.Close()
	bys, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`<div class="g-list" data-v-[a-zA-Z0-9]+>.*[^<]+</div>`)
	res := re.FindAllString(string(bys), 1)
	if len(res) <= 0 {
		fmt.Println("not data.")
		return
	}
	html := res[0]
	re2 := regexp.MustCompile(`<div class="list-item">.*</div>`)
	res2 := re2.FindAllString(html, -1)
	for _, v := range res2 {
		fmt.Println(v)
	}
}
