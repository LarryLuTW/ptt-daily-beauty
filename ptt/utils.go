package ptt

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// fetchPageAmount get latest page number
func fetchPageAmount() int {
	url := "https://www.ptt.cc/bbs/Beauty/index.html"
	doc, _ := goquery.NewDocument(url)
	prevPageSelector := ".wide:nth-child(2)"
	href, _ := doc.Find(prevPageSelector).Attr("href")

	var n int
	fmt.Sscanf(href, "/bbs/Beauty/index%d.html", &n)
	return n
}

// fetchPage fetch all posts in a page
func fetchPage(prefix string, page int) ([]post, error) {
	baseURL := "https://www.ptt.cc/bbs/Beauty/"
	url := fmt.Sprintf("%sindex%d.html", baseURL, page)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	posts := parseDoc2Posts(doc, prefix)
	return posts, nil
}

// sometimes PTT cache search result
// fetchSearchResult use PTT search to get search result
func fetchSearchResult(prefix string, page, recommend int) ([]post, error) {
	// page from 1, 2, ...
	baseURL := "https://www.ptt.cc/bbs/Beauty/search"
	url := fmt.Sprintf("%s?page=%d&q=%s+recommend:%d", baseURL, page, prefix, recommend)
	doc, err := goquery.NewDocument(url)

	if err != nil {
		return nil, err
	}

	posts := parseDoc2Posts(doc, prefix)
	return posts, nil
}

func fetchPreviewImgURL(href string) string {
	// TODO: handle error
	doc, _ := goquery.NewDocument(href)
	imgSelector := `#main-content a[href$=".jpg"],a[href$=".png"],a[href$=".gif"]`
	imgURL, _ := doc.Find(imgSelector).Attr("href")
	return imgURL
}
