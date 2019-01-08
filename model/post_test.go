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
		if b.Href != "https://daily-beauty.xyz/ptt/redirect/M.1543280871.A.39A" {
			t.Errorf("transform url error %s", b.Href)
		}
	}
}
