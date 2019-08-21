package engine

type Item struct {
	Type    string
	Id      string
	Url     string
	Payload interface{}
}

type ParseFunc func(bytes []byte, url string) ParseResult

type Request struct {
	Url       string
	Name      string
	ParseFunc ParseFunc // 编程技巧，函数式编程
}

type ParseResult struct {
	Requests []Request
	Item     *Item
}

type RequestQueue []Request

func (q *RequestQueue) Push(reqs ...Request) {
	*q = append(*q, reqs...)
}

func (q *RequestQueue) Pop() Request {
	request := (*q)[0]
	*q = (*q)[1:]
	return request
}
