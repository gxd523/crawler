package engine

import "crawler/fetcher"

func work(req Request) (*ParseResult, error) {
	if bytes, err := fetcher.Fetch(req.Url); err == nil {
		parseResult := req.ParseFunc(bytes, req.Url)
		return &parseResult, nil
	} else {
		return nil, err
	}
}
