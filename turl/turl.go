package turl

import (
	"encoding/json"
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

func (self Urls) String() string {
	var urls []string
	for i := 0; i < len(self); i++ {
		urls = append(urls, self[i].ExpandedUrl)
	}
	return strings.Join(urls, "\n")
}

func (self Tweet) String() string {
	ret := fmt.Sprintf(
		"(%v)\n%v\n%v%v",
		self.FromUserName,
		self.Text,
		self.Entities.Urls,
		self.Entities.MediaList,
	)
	return strings.Trim(ret, "\n ")
}

// Return tweets for query
//		Example queries:
// 			"Niket"
// 			"Breaking News"
func SearchTweets(query string) (tweets []Tweet, err error) {
	// fetch search response from twitter
	search := "http://search.twitter.com/search.json?%v"
	params := url.Values{
		"q":                []string{query},
		"include_entities": []string{"true"},
		"geocode":          []string{"51.5171,0.1062,100mi"},
	}
	search = fmt.Sprintf(search, params.Encode())

	// Get search response
	response, err := http.Get(search)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// on fail return immediately 
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	// obj is wrapper map around results array of tweets
	var obj SearchResult
	json.Unmarshal(body, &obj)

	tweets = obj.Results

	return
}
