package main

import (
	"flag"
	"fmt"
	"github.com/nexneo/topturls/turl"
	"log"
)

var (
	search string
)

func init() {
	flag.StringVar(&search, "q", "Olympics", "pass query with q")
	flag.Parse()
}

func main() {
	tweets, err := turl.SearchTweets(search)
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := 0; i < len(tweets); i++ {
		fmt.Println(tweets[i], "\n")
	}
}