package engine

import "crawler/model"

type Request struct {
	Url       string
	Name      string
	ParseFunc func([]byte) ParseResult // 编程技巧，函数式编程
}

type ParseResult struct {
	Requests    []Request
	UserProfile *model.UserInfo
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
