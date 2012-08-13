package app

import (
	"code.google.com/p/gorilla/mux"
	"net/http"
)

type Action string

type ShowHandler struct {
	Action
}
type IndexHandler struct {
	Action
}

var (
	showHandler  = ShowHandler{Action("tweet.html")}
	indexHandler = IndexHandler{Action("index.html")}
)

func (handler ShowHandler) ServeHTTP(response http.ResponseWriter,
	req *http.Request) {
	
	//find tweet
	query := req.URL.Query().Get("query")
	id := mux.Vars(req)["id"]
	tweet := search(query).Find(id)

	renderContext := map[string]interface{}{
		"tweet": tweet,
		"query": query,
	}
	handler.Render(response, renderContext)
}

func (handler IndexHandler) ServeHTTP(response http.ResponseWriter,
	req *http.Request) {
	//find tweet
	query := req.FormValue("query")
	if query == "" {
		query = "Olympics"
	}
	tweets := search(query)

	renderContext := map[string]interface{}{
		"tweets": tweets,
		"query":  query,
	}
	handler.Render(response, renderContext)
}
