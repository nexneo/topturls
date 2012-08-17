package app

import (
	"code.google.com/p/gorilla/mux"
	"net/http"
)

func render(t string, response http.ResponseWriter, context interface{}) {
	tpl.Lookup(t).Execute(response, context)
}

func showHandler(response http.ResponseWriter,
	req *http.Request) {

	//find tweet
	query := req.URL.Query().Get("query")
	id := mux.Vars(req)["id"]
	tweet := search(query).Find(id)

	renderContext := map[string]interface{}{
		"tweet": tweet,
		"query": query,
	}
	render("tweet.html", response, renderContext)
}

func indexHandler(response http.ResponseWriter,
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
	render("index.html", response, renderContext)
}
