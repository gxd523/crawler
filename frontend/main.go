package main

import (
	"crawler/frontend/controller"
	"crawler/util"
	"net/http"
)

func main() {
	//execFilePath := "frontend"
	execFilePath := util.GetExecFilePath()
	http.Handle("/", http.FileServer(http.Dir(execFilePath+"/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(execFilePath+"/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
