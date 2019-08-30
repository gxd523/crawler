package worker

import (
	"crawler/engine"
)

type CrawlService struct {
}

func (*CrawlService) Process(req SerializedRequest, result *SerializedParseResult) error {
	request, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	parseResult, err := engine.Work(request)
	if err != nil {
		return err
	}

	*result = SerializeParseResult(parseResult)
	return nil
}
