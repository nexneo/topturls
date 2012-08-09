// Command line twitter search using golang, nothing fancy just learning go.
package main

import (
	//"code.google.com/p/gorilla/mux"
	"flag"
	"github.com/nexneo/topturls/turl"
	"html/template"
	"log"
	"os"
)

var (
	search string
)

func init() {
	flag.StringVar(&search, "q", "Olympics", "pass query with q")
	flag.Parse()
}

func main() {
	tpl, err := template.New("tweet").Parse(tweetTpl)
	tweets, err := turl.SearchTweets(search)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, tweet := range tweets {
		tpl.Execute(os.Stdout, tweet)
	}
}

const (
	tweetTpl = `
({{.FromUserName}})
{{.Text}}
{{range .Entities.Urls}}{{.ExpandedUrl}}{{end}}
`
)
