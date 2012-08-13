package app

import (
	"github.com/nexneo/topturls/turl"
	"log"
)

var (
	Searches = map[string]turl.Tweets{}
)

func search(query string) turl.Tweets {
	var tweets turl.Tweets
	var ok bool
	var err error

	if tweets, ok = Searches[query]; !ok {
		tweets, err = turl.SearchTweets(query)

		if err != nil {
			log.Fatal(err)
		}

		Searches[query] = tweets
	}
	return tweets
}
