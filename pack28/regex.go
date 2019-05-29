package main

import (
	"fmt"
	"regexp"
)

func main()  {
	s := `

fdsfdsdemofdsfdemofsdfsddemo demo@qq.com fsdfsd
fdsfdsdemofdsfdemofsdfsddemo demoaaa@qq.com fsdfsd
fdsfdsdemofdsfdemofsdfsddemo democcc@qq.com fsdfsd
fdsfdsdemofdsfdemofsdfsddemo demoggg@qq.com fsdfsd
fdsfdsdemofdsfdemofsdfsddemo demoppp@qq.com fsdfsd
fdsfdsdemofdsfdemofsdfsddemo demozzz@qq.com fsdfsd


`
	re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\.\w+`)
	rs := re.FindAllString(s, -1)
	for _, mail := range rs {
		fmt.Println(mail)
	}
}
