package duplicate

import (
	"log"
)

type RepeatService struct {
	fetchingUrlMap map[string]bool
}

func (service *RepeatService) IsDuplicateUrl(url string, result *bool) error {
	if service.fetchingUrlMap == nil {
		service.fetchingUrlMap = make(map[string]bool)
	}
	if isVisited := service.fetchingUrlMap[url]; isVisited {
		*result = true
		log.Printf("duplicate url->%s\n", url)
	} else {
		service.fetchingUrlMap[url] = true
		*result = false
	}
	return nil
}
