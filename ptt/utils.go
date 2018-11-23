package ptt

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/vjeantet/jodaTime"
)

// FIXME: sometimes PTT cache search result
func fetchSearchResult(prefix string, page, recommend int) ([]post, error) {
	baseURL := "https://www.ptt.cc/bbs/Beauty/search"
	url := fmt.Sprintf("%s?page=%d&q=[+recommend:%d", baseURL, page, recommend)
	doc, err := goquery.NewDocument(url)

	if err != nil {
		return nil, err
	}

	posts := make([]post, 0, 20)
	doc.Find(".r-ent").Each(func(i int, el *goquery.Selection) {
		nVoteText := el.Find(".hl").Text()
		nVote := parseNVote(nVoteText)

		titleEl := el.Find(".title > a")
		title := titleEl.Text()

		if !strings.HasPrefix(title, prefix) {
			return
		}

		hrefText, _ := titleEl.Attr("href")
		href := "https://www.ptt.cc" + hrefText

		currentYear := time.Now().Year()
		dateText := fmt.Sprintf("%d/%s", currentYear, el.Find(".meta .date").Text())
		date, _ := jodaTime.ParseInLocation("YYYY/MM/dd", dateText, "Asia/Taipei")

		p := post{
			title: title,
			href:  href,
			nVote: nVote,
			date:  date,
		}

		posts = append(posts, p)
	})

	return posts, nil
}

func fetchPreviewImgURL(href string) string {
	doc, _ := goquery.NewDocument(href)
	imgSelector := `#main-content a[href$=".jpg"],a[href$=".png"],a[href$=".gif"]`
	imgURL, _ := doc.Find(imgSelector).Attr("href")
	return imgURL
}
