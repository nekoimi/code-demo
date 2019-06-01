package parser

import (
	"../../conf"
	"../../engine"
	"../../fetcher"
	"errors"
	"regexp"
)

const clearHtmlRegex = `\s{2,}`
const goodListRegex = `<li data\-index="[0-9]+" data\-goods\-id="[0-9]+" class="ga\-list g\-item"[^>]*>.*?<\/li>`
const goodListItemOneRegex = `<a href="(\/category\-shafa\/goods\-[0-9]*\.html\?page=[0-9]+&index=[0-9]+)" target="_blank">.*<img class="d\-img" src="//.*"[^>]*data\-src="(.*?)"[^>]*data\-wide\-src="//.*"[^>]*alt="(.*?)图片"[^>]*><\/a>`
const goodListItemTwoRegex = `<a href="(\/category\-shafa\/goods\-[0-9]*\.html\?page=[0-9]+&index=[0-9]+)" target="_blank">.*<img class="d\-img" src="(.*?)"[^>]*alt="(.*?)图片"[^>]*><\/a>`

// 商品列表解析器
func ParseGoodList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(goodListRegex)
	html := fetcher.RegexpReplaceEmpty(content, clearHtmlRegex)
	res := re.FindAllString(html, -1)
	result := engine.ParseResult{}
	for _, match := range res {
		url, itemName, itemErr := parseGoodListItem(match)
		if itemErr != nil {
			continue
		}
		result.Items = append(result.Items, itemName)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParseFunc: ParseGoodDetails,
		})
	}
	return result
}

// 商品列表Item解析器
func parseGoodListItem(html string) (string, string,error) {
	var resItem []string
	reOne := regexp.MustCompile(goodListItemOneRegex)
	findRes := reOne.FindString(html)
	if findRes == "" {
		reTwo := regexp.MustCompile(goodListItemTwoRegex)
		resItem = reTwo.FindStringSubmatch(html)
	} else {
		resItem = reOne.FindStringSubmatch(html)
	}
	if len(resItem) < 3 {
		return "", "", errors.New("Goods item not found.")
	}

	return conf.RootUrl + resItem[1], string(resItem[3]), nil
}
