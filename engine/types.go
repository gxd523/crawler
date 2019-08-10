package engine

type Request struct {
	Url       string
	Name      string
	ParseFunc func([]byte) []Request
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
