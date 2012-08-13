package turl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	
	Query string
}

type Tweets []Tweet

type SearchResult struct {
	Page       uint    `json:"page"`
	Query      string  `json:"query"`
	Tweets    Tweets `json:"results"`
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
func SearchTweets(query string) (Tweets, error) {
	var err error
	// fetch search response from twitter
	searchURL := "http://search.twitter.com/search.json?%s"
	params := url.Values{
		"q":                {query},
		"include_entities": {"true"},
	}
	searchURL = fmt.Sprintf(searchURL, params.Encode())

	// Get search response
	response, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		status := fmt.Sprintf(": %s", response.Status)
		return nil, errors.New(status)
	}

	// on fail return immediately 
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// obj is wrapper map around results array of tweets
	var result SearchResult
	var tweets Tweets

	json.Unmarshal(body, &result)
	
	for _, tweet := range result.Tweets {
		tweet.Query = result.Query
		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
