package worker

import (
	"crawler/distributed/config"
	"crawler/engine"
	"crawler/parser"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	FuncName string
	Args     interface{}
}

type SerializedRequest struct {
	Url    string
	Name   string
	Parser SerializedParser
}

type SerializedParseResult struct {
	Requests []SerializedRequest
	Item     *engine.Item
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.FuncName { // 可以讲方法名放进map集合
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseUserInfo:
		if args, ok := p.Args.([]interface{}); ok {
			username := args[0].(string)
			gender := args[1].(string)
			return &parser.UserInfoParser{
				Username: username,
				Gender:   gender,
			}, nil
		} else {
			return nil, fmt.Errorf("error: invalid args: %v\n", p.Args)
		}
	default:
		return nil, errors.New("error: deserialize parser failed")
	}
}

func SerializeRequest(req engine.Request) SerializedRequest {
	funcName, args := req.Parser.Serialize()
	return SerializedRequest{
		Url:  req.Url,
		Name: req.Name,
		Parser: SerializedParser{
			FuncName: funcName,
			Args:     args,
		},
	}
}

func DeserializeRequest(req SerializedRequest) (engine.Request, error) {
	deserializedParser, err := deserializeParser(req.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    req.Url,
		Name:   req.Name,
		Parser: deserializedParser,
	}, nil
}

func SerializeParseResult(result *engine.ParseResult) SerializedParseResult {
	parseResult := SerializedParseResult{
		Item: result.Item,
	}
	for _, req := range result.Requests {
		parseResult.Requests = append(parseResult.Requests, SerializeRequest(req))
	}
	return parseResult
}

func DeserializeParseResult(result SerializedParseResult) engine.ParseResult {
	parseResult := engine.ParseResult{
		Item: result.Item,
	}
	for _, req := range result.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error: deserialized request failed")
			continue
		}
		parseResult.Requests = append(parseResult.Requests, request)
	}
	return parseResult
}
