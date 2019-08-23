package view

import (
	"crawler/engine"
	"crawler/frontend/model"
	model2 "crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	searchResultView := CreateSearchResultView("template.html")
	page := model.SearchResult{
		Hits:     20,
		Start:    1,
		Query:    "male",
		PrevFrom: 1,
		NextFrom: 28,
	}

	item := engine.Item{
		Type: "",
		Id:   "",
		Url:  "",
		Payload: model2.UserInfo{
			Name:       "郭晓东",
			Gender:     "male",
			Age:        28,
			Height:     176,
			Income:     "1-2W",
			Marriage:   "married",
			Education:  "college",
			Occupation: "Hangzhou",
			Birthplace: "Xinjiang",
			House:      "Yes",
			Car:        "No",
			Xinzuo:     "Gemini",
		},
	}

	for i := 0; i < int(page.Hits); i++ {
		page.Items = append(page.Items, item)
	}

	file, err := os.Create("test.html")
	err = searchResultView.Render(file, page)
	if err != nil {
		panic(err)
	}
}
