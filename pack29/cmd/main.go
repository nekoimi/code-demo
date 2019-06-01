package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main()  {
	response, err := http.Get("")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("status code ===> ", response.StatusCode)
		return
	}

	//e := determineEncoding(response.Body)
	//utf8Reader := transform.NewReader(response.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//html := string(all)
	//err = ioutil.WriteFile("./html.html", all, 0777)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(html)
	re := regexp.MustCompile(`<a href="\/category-keting\/list-b[0-9]+\/\?from=b" title="(.*)" class="abs"><span class="txt"[^>]*>[^<]+<\/span><\/a>`)
	res := re.FindAllSubmatch(all, -1)

	var brands []string
	for _, match := range res {
		brands = append(brands, string(match[1]))
	}

	url := "https://new-retail.yoyohr.com/weapi/v1/brand"

	for _, brand := range brands {
		br := make(map[string]string)
		br["name"] = brand
		fmt.Println("Add ===> ", brand)
		reader, err := json.Marshal(br)
		if err != nil {
			fmt.Println("json err.   ", reader)
			continue
		}
		request, _ := http.NewRequest("POST", url, bytes.NewReader(reader))
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
		request.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL25ldy1yZXRhaWwueW95b2hyLmNvbS93ZWFwaS92MS9icmFuZHMiLCJpYXQiOjE1NTkzMTEwNDAsImV4cCI6MTU1OTMyMTQ4NSwibmJmIjoxNTU5MzE3ODg1LCJqdGkiOiJOelZnSjVTS1NyT2FsOTdzIiwic3ViIjoiNmIxMjM3YmI2NWZjN2RhMDAxMTQzYmNhNjFhYjVjMTgiLCJwcnYiOiJjODI5MjIzODM1ZDExMTM4ZjA4YWNlNTZmZmE2NjI4YmMyNjgzY2I1In0.T8AKRolM_bAGCmQ6vcegLHtHktFHZend01kBt1HeRWE")
		client := http.Client{}
		res, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.StatusCode)
	}

}

//func determineEncoding(r io.Reader) encoding.Encoding {
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		panic(err)
//	}
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//	return e
//}
