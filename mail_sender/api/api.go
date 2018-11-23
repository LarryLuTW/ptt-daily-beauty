package api

import (
	"fmt"
	"main/model"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/vjeantet/jodaTime"
)

const url = "https://us-central1-daily-beauty-209105.cloudfunctions.net/getDailyBeauties"

// parseNVote parses vote text to int
// "50" => 50, "爆" => 100
// there is no need to handle nVote <= 0
// because they are filterer out when searching
func parseNVote(nVoteText string) int {
	if nVoteText == "爆" {
		return 100
	}
	nVote, _ := strconv.Atoi(nVoteText)
	return nVote
}

type post struct {
	title string
	href  string
	nVote int
	date  time.Time
}

func isToday(t time.Time) bool {
	loc, _ := time.LoadLocation("Asia/Taipei")
	current := time.Now().In(loc)
	return t.YearDay() == current.YearDay()
}

func isYesterday(t time.Time) bool {
	// FIXME: 跨年 1/1 < 12/31
	loc, _ := time.LoadLocation("Asia/Taipei")
	current := time.Now().In(loc)
	return t.YearDay() == current.YearDay()-1
}

func isBeforeYesterday(t time.Time) bool {
	// FIXME: 跨年 1/1 < 12/31
	loc, _ := time.LoadLocation("Asia/Taipei")
	current := time.Now().In(loc)
	return t.YearDay() < current.YearDay()-1
}

// TODO: function too big
func fetchYesterPosts() []post {
	baseURL := "https://www.ptt.cc/bbs/Beauty/search"
	keyword := "["
	posts := make([]post, 0, 20)

	page := 1

	for {

		url := fmt.Sprintf("%s?page=%d&q=%s+recommend:1", baseURL, page, keyword)

		doc, err := goquery.NewDocument(url)

		if err != nil {
			panic(err)
		}

		doc.Find(".r-ent").Each(func(i int, el *goquery.Selection) {
			nVoteText := el.Find(".hl").Text()
			nVote := parseNVote(nVoteText)

			titleEl := el.Find(".title > a")
			title := titleEl.Text()

			if !strings.HasPrefix(title, "[正妹]") {
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

		oldestDate := posts[len(posts)-1].date
		if isBeforeYesterday(oldestDate) {
			break
		}
		page++
	}

	// filter yesterday post
	yesterdayPosts := make([]post, 0, 10)
	for _, p := range posts {
		if isYesterday(p.date) {
			yesterdayPosts = append(yesterdayPosts, p)
		}
	}

	return yesterdayPosts
}

// TODO: rename
func getChampions(posts []post) []model.Beauty {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].nVote > posts[j].nVote
	})

	champions := posts[:3]
	beauties := make([]model.Beauty, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	for i, p := range champions {
		go func(i int, p post) {
			defer wg.Done()
			doc, _ := goquery.NewDocument(p.href)
			imgSelector := `#main-content a[href$=".jpg"],a[href$=".png"],a[href$=".gif"]`
			imgURL, _ := doc.Find(imgSelector).Attr("href")
			beauties[i] = model.Beauty{
				NVote:      p.nVote,
				Title:      p.title,
				Href:       p.href,
				PreviewImg: imgURL,
			}
		}(i, p)
	}

	wg.Wait()

	beauties[0].Rank = "一"
	beauties[1].Rank = "二"
	beauties[2].Rank = "三"

	return beauties
}

// FetchBeauties send a request to get beauties from getDailyBeauties api
func FetchBeauties() ([]model.Beauty, error) {
	posts := fetchYesterPosts()
	beauties := getChampions(posts)

	return beauties, nil
}
