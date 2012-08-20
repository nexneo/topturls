package turl

import (
	"testing"
)

func TestMedia(t *testing.T) {
	tweets, err := SearchTweets("bachhan")
	if err != nil{
		t.Error(err)
	}
	if len(tweets) > 1 {
		t.Errorf("%d", len(tweets))
	}
}
