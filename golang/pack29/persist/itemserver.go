package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for  {
			item := <- out
			log.Printf("Item Count #%d : %v", itemCount, item)
			saveItem(item)
			itemCount++
		}
	}()
	return out
}

func saveItem(item interface{})  {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	_, indexErr := client.Index().Index("spider").
		Type("goods").
		BodyJson(item).
		Do(context.Background())
	if indexErr != nil {
		log.Println(indexErr)
	}
}
