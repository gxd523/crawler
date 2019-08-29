package persist

import (
	"crawler/engine"
	"gopkg.in/olivere/elastic.v6"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := Save(item, s.Client, s.Index)
	if err == nil {
		*result = "ok"
	}
	return err
}
