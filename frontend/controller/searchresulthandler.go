package controller

import (
	"context"
	"crawler/engine"
	"crawler/frontend/model"
	"crawler/frontend/view"
	"gopkg.in/olivere/elastic.v6"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// http://localhost:8888/search?q=男 已购房&from=20&Payload.Age:(<30)
func (handler SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := handler.getSearchResult(q, from)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = handler.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (handler SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var searchResult model.SearchResult

	search := handler.client.Search("dating_userinfo")
	if q != "" {
		search.Query(elastic.NewQueryStringQuery(rewriteQueryString(q)))
	}
	elasticSearchResult, err := search.From(from).Do(context.Background())

	if err != nil {
		return searchResult, nil
	}
	searchResult.Query = q
	searchResult.Hits = elasticSearchResult.TotalHits()
	searchResult.Start = from
	var items []engine.Item
	for _, v := range elasticSearchResult.Each(reflect.TypeOf(engine.Item{})) {
		item := v.(engine.Item)
		items = append(items, item)
	}
	searchResult.Items = items

	searchResult.PrevFrom = searchResult.Start - len(searchResult.Items)
	searchResult.NextFrom = searchResult.Start + len(searchResult.Items)

	return searchResult, nil
}

func CreateSearchResultHandler(templateFileName string) SearchResultHandler {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(templateFileName),
		client: client,
	}
}

func rewriteQueryString(q string) string {
	compile := regexp.MustCompile(`([A-Za-z]*):`)
	return compile.ReplaceAllString(q, "Payload.$1:")
}
