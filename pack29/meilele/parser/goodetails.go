package parser

import (
	"../../engine"
	"../../fetcher"
	"regexp"
)

const goodDetailImagesRegex = `<div class="img" data\-source="handledpictures" style="margin\-bottom:[0-9]+px;">[^<]*<img src="//image\.meilele\.com\/images\/blank\.gif"[^>]*data\-src="//(.+\.(jpg|png|bum|gif|jpeg|ico))"`

// 商品详情解析器
func ParseGoodDetails(content []byte) engine.ParseResult {
	html := fetcher.RegexpReplaceEmpty(content, clearHtmlRegex)

	detailRe := regexp.MustCompile(goodListItemTwoRegex)
	results := detailRe.FindAllStringSubmatch(html, -1)
	var images []string
	for _, image := range results {
		images = append(images, image[1])
	}
	if len(images) < 3 {
		return engine.ParseResult{}
	}

}
