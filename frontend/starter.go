package main

import (
	"github.com/iralance/go-clawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
