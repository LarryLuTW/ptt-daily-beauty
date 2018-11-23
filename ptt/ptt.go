package ptt

import (
	"fmt"
	"main/model"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/vjeantet/jodaTime"
)

const url = "https://us-central1-daily-beauty-209105.cloudfunctions.net/getDailyBeauties"

type post struct {
	title string
	href  string
	nVote int
	date  time.Time
}

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

// TODO: return error
func fetchYesterdayPosts() []post {
	prefix := "[正妹]"
	posts := make([]post, 0, 20)

	page := 1
	for {
		ps, _ := fetchSearchResult(prefix, page, 1)
		posts = append(posts, ps...)

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

	// [正妹] 大橋未久 -> 大橋未久
	for i := range beauties {
		beauties[i].Title = beauties[i].Title[9:]
	}

	return beauties
}

// FetchBeauties send a request to get beauties from getDailyBeauties api
func FetchBeauties() ([]model.Beauty, error) {
	posts := fetchYesterdayPosts()
	beauties := getChampions(posts)

	return beauties, nil
}
