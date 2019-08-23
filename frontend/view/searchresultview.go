package view

import (
	"crawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}

func CreateSearchResultView(fileName string) SearchResultView {
	return SearchResultView{template: template.Must(template.ParseFiles(fileName))}
}
