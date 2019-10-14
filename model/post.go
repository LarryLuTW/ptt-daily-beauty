package model

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Post is a corresponding to a post on ptt
type Post struct {
	Title string
	Href  string
	NVote int
	Date  time.Time
}

// fetchPreviewImg get the preview image of a post
func fetchPreviewImg(p *Post) string {
	// TODO: handle error
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", p.Href, nil)
	req.Header.Set("Cookie", "over18=1")

	res, _ := client.Do(req)
	doc, _ := goquery.NewDocumentFromResponse(res)

	imgSelector := `#main-content a[href$=".jpg"],a[href$=".png"],a[href$=".gif"]`
	imgURL, _ := doc.Find(imgSelector).Attr("href")
	return imgURL
}

// fetchImageAmount get the amount of images in a post
func fetchImageAmount(p *Post) int {
	// TODO: handle error
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", p.Href, nil)
	req.Header.Set("Cookie", "over18=1")

	res, _ := client.Do(req)
	doc, _ := goquery.NewDocumentFromResponse(res)

	doc.Find("div.push").Each(func(i int, s *goquery.Selection) {
		// remove push comment
		s.Remove()
	})
	imgSelector := `#main-content a[href$=".jpg"],a[href$=".png"],a[href$=".gif"]`
	nImage := doc.Find(imgSelector).Size()
	return nImage
}

// "[正妹] 大橋未久" -> "大橋未久"
func trimTitlePrefix(title string) string {
	return strings.TrimPrefix(title, "[正妹] ")
}

// transform https://www.ptt.cc/bbs/Beauty/M.1543991133.A.1A1.html
// to https://daily-beauty.xyz/ptt/redirect/M.1543991133.A.1A1
func transformURL(pttURL string) string {
	var articleID string
	fmt.Sscanf(pttURL, "https://www.ptt.cc/bbs/Beauty/%18s.html", &articleID)
	return fmt.Sprintf("https://daily-beauty.xyz/ptt/redirect/%s", articleID)
}

// ToBeauty transform a Post to a Beauty
func (p *Post) ToBeauty() Beauty {
	previewImg := fetchPreviewImg(p)
	nImage := fetchImageAmount(p)
	return Beauty{
		NVote:      p.NVote,
		NImage:     nImage,
		Title:      trimTitlePrefix(p.Title),
		Href:       transformURL(p.Href),
		PreviewImg: previewImg,
	}
}
