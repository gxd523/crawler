package persist

import (
	"context"
	"crawler/model"
	"crawler/src/github.com/olivere/elastic"
	"encoding/json"
	"testing"
)

func TestSave(t *testing.T) {
	expectedUserInfo := model.UserInfo{
		Name:       "郭晓东",
		Gender:     "male",
		Age:        28,
		Height:     176,
		Weight:     60,
		Income:     "1-2W",
		Marriage:   "married",
		Education:  "college",
		Occupation: "Hangzhou",
		Birthplace: "Xinjiang",
		House:      "Yes",
		Car:        "No",
		Xinzuo:     "Gemini",
	}
	id, err := save(expectedUserInfo)
	if err != nil {
		panic(err)
	}

	// TODO Try to start up elastic search(here using docker go client.)
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	result, err := client.Get().
		Index("dating_userinfo").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actualUserInfo model.UserInfo
	err = json.Unmarshal(result.Source, &actualUserInfo)
	if err != nil {
		panic(err)
	}

	if expectedUserInfo != actualUserInfo {
		t.Errorf("got %+v, expected %v", actualUserInfo, expectedUserInfo)
	}
}
