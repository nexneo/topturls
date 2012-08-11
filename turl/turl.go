package turl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Media struct {
	DisplayUrl    string `json:"display_url"`
	ExpandedUrl   string `json:"expanded_url"`
	IdStr         string `json:"id_str"`
	MediaUrl      string `json:"media_url"`
	MediaUrlHttps string `json:"media_url_https"`
	Type          string `json:"type"`
	Url           string `json:"url"`
}

type Urls []Media

type Mention struct {
	IdStr      string `json:"id_str"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type Tweet struct {
	Entities struct {
		MediaList Urls      `json:"media"`
		Urls      Urls      `json:"urls"`
		Mentions  []Mention `json:"user_mentions"`
	}
	Metadata struct {
		ResultType string `json:"result_type"`
	}
	CreatedAt       string `json:"created_at"`
	FromUser        string `json:"from_user"`
	FromUserId      uint64 `json:"from_user_id"`
	FromUserIdStr   string `json:"from_user_id_str"`
	FromUserName    string `json:"from_user_name"`
	Id              uint64 `json:"id"`
	IdStr           string `json:"id_str"`
	ProfileImageUrl string `json:"profile_image_url"`
	Source          string `json:"source"`
	Text            string `json:"text"`
	ToUserIdStr     string `json:"to_user_id_str"`
}

type Tweets []Tweet

type SearchResult struct {
	Page       uint    `json:"page"`
	Query      string  `json:"query"`
	Results    []Tweet `json:"results"`
	SinceIdStr string  `json:"since_id_str"`
}

func (tweets Tweets) Find(idstr string) (tweet Tweet) {
	for _, tweet = range tweets {
		if tweet.IdStr == idstr {
			break
		}
	}
	return
}

// Return tweets for query
func SearchTweets(results chan Tweet, query string) {
	var err error
	// fetch search response from twitter
	search := "http://search.twitter.com/search.json?%s"
	params := url.Values{
		"q":                {query},
		"include_entities": {"true"},
	}
	search = fmt.Sprintf(search, params.Encode())

	// Get search response
	response, err := http.Get(search)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		status := fmt.Sprintf(": %s", response.Status)
		log.Fatal(errors.New(status))
	}

	// on fail return immediately 
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// obj is wrapper map around results array of tweets
	var obj SearchResult
	json.Unmarshal(body, &obj)

	for _, tweet := range obj.Results {
		results <- tweet
	}
	close(results)
}
