package persist

import (
	"../engine"
	"fmt"
	"testing"
)

func TestItemSave(t *testing.T) {
	goods := engine.Goods{
		Title: "goods title",
		Price: 233,
		Banners: []string {""},
		Brand: "brand name",
		Details: []string{""},
	}
	fmt.Println(goods)
	saveItem(goods)
}
