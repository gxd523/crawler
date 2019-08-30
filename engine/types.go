package engine

type Item struct {
	Type    string
	Id      string
	Url     string
	Payload interface{}
}

type ParseFunc func(bytes []byte, url string) *ParseResult

type Parser interface {
	Parse(bytes []byte, url string) *ParseResult

	Serialize() (funcName string, args interface{})
}

type FuncParser struct {
	ParseFunc ParseFunc
	FuncName  string
}

func (f *FuncParser) Parse(bytes []byte, url string) *ParseResult {
	return f.ParseFunc(bytes, url)
}

func (f *FuncParser) Serialize() (funcName string, args interface{}) {
	return f.FuncName, nil
}

func NewFuncParser(parseFunc ParseFunc, funcName string) *FuncParser {
	return &FuncParser{
		ParseFunc: parseFunc,
		FuncName:  funcName,
	}
}

type Request struct {
	Url    string
	Name   string
	Parser Parser
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
