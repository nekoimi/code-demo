package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type Goods struct {
	Title   string   `json:"title"`
	Price   float64  `json:"price"`
	Banners []string `json:"banners"`
	Brand   string   `json:"brand"`
	Details []string
}

func NilParseResult([]byte) ParseResult {
	return ParseResult{}
}
