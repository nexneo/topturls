package turl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

type SearchResult struct {
	Page       uint    `json:"page"`
	Query      string  `json:"query"`
	Results    []Tweet `json:"results"`
	SinceIdStr string  `json:"since_id_str"`
}

func (urls Urls) String() string {
	var ret []string
	for _, u := range urls {
		ret = append(ret, u.ExpandedUrl)
	}
	return strings.Join(ret, "\n")
}

func (tweet Tweet) String() string {
	ret := fmt.Sprintf(
		"(%v)\n%v\n%v%v",
		tweet.FromUserName,
		tweet.Text,
		tweet.Entities.Urls,
		tweet.Entities.MediaList,
	)
	return strings.Trim(ret, "\n ")
}

// Return tweets for query
//		Example queries:
// 			"Niket"
// 			"Breaking News"
func SearchTweets(query string) ([]Tweet, error) {
	var tweets []Tweet
	var err error
	// fetch search response from twitter
	search := "http://search.twitter.com/searcd.json?%s"
	params := url.Values{
		"q":                {query},
		"include_entities": {"true"},
		"geocode":          {"37.33182,122.03118,100mi"},
	}
	search = fmt.Sprintf(search, params.Encode())

	// Get search response
	response, err := http.Get(search)
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
	var obj SearchResult
	json.Unmarshal(body, &obj)

	tweets = obj.Results

	return tweets, nil
}
