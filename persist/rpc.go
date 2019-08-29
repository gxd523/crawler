package persist

import (
	"crawler/engine"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

type ItemSaverService struct {
	Client    *elastic.Client
	Index     string
	itemCount int
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := Save(item, s.Client, s.Index)
	if err == nil {
		*result = "ok"
		s.itemCount++
		log.Printf("%d...%+v\n", s.itemCount, item)
	} else {
		log.Printf("Error saving %d Item %+v: %v\n", s.itemCount, item, err)
	}
	return err
}
