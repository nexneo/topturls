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

func (self Urls) String() string {
	var urls []string
	for i := 0; i < len(self); i++ {
		urls = append(urls, self[i].ExpandedUrl)
	}
	return strings.Join(urls, "\n")
}

type Mention struct {
	IdStr      string `json:"id_str"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type Tweet struct {
	Entities struct {
		MediaList []Media   `json:"media"`
		Urls      Urls      `json:"urls"`
		Mentions  []Mention `json:"user_mentions"`
	}
	Metadata struct {
		ResultType string `json:"result_type"`
	}
	CreatedAt       string `json:"created_at"`
	FromUser        string `json:"from_user"`
	FromUserId      int    `json:"from_user_id"`
	FromUserIdStr   string `json:"from_user_id_str"`
	FromUserName    string `json:"from_user_name"`
	IdStr           string `json:"id_str"`
	IsoLanguageCode string `json:"iso_language_code"`
	Text            string `json:"text"`
}

func (self Tweet) String() string {
	ret := fmt.Sprintf(
		"(%v)\n%v\n%v",
		self.FromUserName,
		self.Text,
		self.Entities.Urls,
	)
	return strings.Trim(ret, "\n ")
}

type SearchResult struct {
	MaxIdStr       string  `json:"max_id_str"`
	NextPage       string  `json:"next_page"`
	Page           int     `json:"page"`
	Query          string  `json:"query"`
	RefreshUrl     string  `json:"refresh_url"`
	Results        []Tweet `json:"results"`
	ResultsPerPage int     `json:"results_per_page"`
	SinceIdStr     string  `json:"since_id_str"`
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
