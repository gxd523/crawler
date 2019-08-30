package engine

import "crawler/fetcher"

func Work(req Request) (*ParseResult, error) {
	if bytes, err := fetcher.Fetch(req.Url); err == nil {
		return req.Parser.Parse(bytes, req.Url), nil
	} else {
		return nil, err
	}
}
