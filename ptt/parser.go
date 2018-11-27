package ptt

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/vjeantet/jodaTime"
)

// parseNVote parses vote text to int
// "50" => 50, "爆" => 100
// "" => 0
// "X7" => -1
// there is no need to handle nVote <= 0
// because they are filterer out when searching
func parseNVote(nVoteText string) int {
	if nVoteText == "爆" {
		return 100
	}
	if nVoteText == "" {
		return 0
	}
	if strings.HasPrefix(nVoteText, "X") {
		return -1
	}
	nVote, _ := strconv.Atoi(nVoteText)
	return nVote
}

// [正妹] 大橋未久 -> 大橋未久
func trimTitlePrefix(title string) string {
	return strings.TrimPrefix(title, "[正妹] ")
}

func parseDoc2Posts(doc *goquery.Document, prefix string) []post {
	// TODO: remove 置頂文
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
	return posts
}
