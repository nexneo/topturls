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
	results := make(chan turl.Tweet, 1)
	go turl.SearchTweets(results, search)
	tpl, err := template.New("tweet").Parse(tweetTpl)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Searching...", search)
	for tweet := range results {
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
