package api

import (
	"strings"
	"testing"
)

func TestFetchPageAmount(t *testing.T) {
	nPage, err := FetchPageAmount()
	if nPage < 2733 {
		t.Errorf("nPage is wrong")
	}
	if err != nil {
		t.Error(nil)
	}
}

func TestFetchPage(t *testing.T) {
	prefix := "["
	posts, err := FetchPage(prefix, 2733)
	if err != nil {
		t.Error(err)
	}
	if len(posts) == 0 {
		t.Errorf("posts should NOT be empty")
	}
	if !strings.HasPrefix(posts[0].Title, prefix) {
		t.Errorf("posts should be prefixed with %s", prefix)
	}
}

func TestSearch(t *testing.T) {
	prefix := "["
	posts, err := Search(prefix, 1, 10)
	if err != nil {
		t.Error(err)
	}
	if len(posts) == 0 {
		t.Errorf("len(posts) should be 20")
	}
	if posts[0].NVote < 10 {
		t.Errorf("post.NVote should >= 10")
	}
}
