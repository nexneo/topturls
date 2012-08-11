package app

import (
	"code.google.com/p/gorilla/mux"
	"github.com/nexneo/topturls/turl"
	"go/build"
	"html/template"
	"log"
	"net/http"
)

const basePkg = "github.com/nexneo/topturls"

type Action string

type IndexHandler struct {
	A Action
}
type TweetHandler struct {
	A Action
}
type SearchHandler struct {
	A Action
}

var (
	indexH = IndexHandler{Action("index.html")}
	tweetH = TweetHandler{Action("tweet.html")}
	searcH = SearchHandler{Action("index.html")}
	tpl    *template.Template
	tweets turl.Tweets
	router *mux.Router
)

func init() {
	p, err := build.Default.Import(basePkg, "", build.FindOnly)
	if err != nil {
		log.Fatalf("Couldn't find staic: %v", err)
	}
	rootPath := p.Dir
	templatePath := rootPath + "/templates/*.html"
	tpl = template.Must(template.ParseGlob(templatePath))

	router = mux.NewRouter()
	router.Handle("/", indexH)
	router.Handle("/tweet/{id}", tweetH)
	router.Handle("/search", searcH)

	http.Handle("/", router)
	http.Handle("/public/", http.FileServer(http.Dir(rootPath)))
}

func (handler IndexHandler) ServeHTTP(response http.ResponseWriter,
	req *http.Request) {

	results := make(chan turl.Tweet, 100)
	search := "Olympics"
	go turl.SearchTweets(results, search)

	tweets = make(turl.Tweets, 0)
	for tweet := range results {
		tweets = append(tweets, tweet)
	}

	handler.A.Render(response, tweets)
}

func (handler TweetHandler) ServeHTTP(response http.ResponseWriter,
	req *http.Request) {
	//find tweet
	tweet := tweets.Find(mux.Vars(req)["id"])
	handler.A.Render(response, tweet)
}

func (handler SearchHandler) ServeHTTP(response http.ResponseWriter,
	req *http.Request) {
	//find tweet
	results := make(chan turl.Tweet, 100)
	search := req.FormValue("query")
	go turl.SearchTweets(results, search)

	tweets = make(turl.Tweets, 0)
	for tweet := range results {
		tweets = append(tweets, tweet)
	}

	handler.A.Render(response, tweets)
}

func (t Action) Render(response http.ResponseWriter, context interface{}) {
	tpl.Lookup("head.html").Execute(response, nil)
	tpl.Lookup(string(t)).Execute(response, context)
	tpl.Lookup("tail.html").Execute(response, nil)
}

func Start() {
	//start http server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
