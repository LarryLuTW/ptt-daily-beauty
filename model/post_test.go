package model

import (
	"testing"
	"time"
)

func TestPostToBeauty(t *testing.T) {
	{
		p := Post{
			Title: "[正妹] 覺得還不錯",
			Href:  "https://www.ptt.cc/bbs/Beauty/M.1543280871.A.39A.html",
			NVote: 50,
			Date:  time.Now(),
		}
		b := p.ToBeauty()
		if b.PreviewImg != "https://imgur.com/30XW9qD.jpg" {
			t.Error("preview image error")
		}
		if b.Title != "覺得還不錯" {
			t.Error("trim title error")
		}
	}

	{
		p := Post{
			Title: "[正妹] 這個笑容",
			Href:  "https://www.ptt.cc/bbs/Beauty/M.1479381920.A.D30.html",
			NVote: 100,
			Date:  time.Now(),
		}
		b := p.ToBeauty()
		if b.PreviewImg != "http://imgur.com/8bsl72C.jpg" {
			t.Error("preview image error")
		}
		if b.Title != "這個笑容" {
			t.Error("trim title error")
		}
	}
}
