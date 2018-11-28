package ptt

import (
	"github.com/PuerkitoBio/goquery"
)

func fetchPreviewImgURL(href string) string {
	// TODO: handle error
	doc, _ := goquery.NewDocument(href)
	imgSelector := `#main-content a[href$=".jpg"],a[href$=".png"],a[href$=".gif"]`
	imgURL, _ := doc.Find(imgSelector).Attr("href")
	return imgURL
}
