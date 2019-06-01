package main

import (
	"../engine"
	"../meilele/parser"
)

func main()  {
	engine.Run(engine.Request{
		Url: "https://www.meilele.com/category-keting/list-p1/?from=page",
		ParseFunc: parser.ParseGoodList,
	})
}