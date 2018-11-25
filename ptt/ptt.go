package ptt

import (
	"main/model"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type post struct {
	title string
	href  string
	nVote int
	date  time.Time
}

func fetchYesterdayPosts() ([]post, error) {
	prefix := "[正妹]"
	recentPosts := make([]post, 0, 20)

	// get recent posts
	page := 1
	for ; ; page++ {
		posts, err := fetchSearchResult(prefix, page, 1)

		if err != nil {
			return nil, err
		}

		recentPosts = append(recentPosts, posts...)
		oldestDate := recentPosts[len(recentPosts)-1].date
		if isBeforeYesterday(oldestDate) {
			break
		}
	}

	// filter yesterday post
	yesterdayPosts := make([]post, 0, 10)
	for _, p := range recentPosts {
		if isYesterday(p.date) {
			yesterdayPosts = append(yesterdayPosts, p)
		}
	}

	return yesterdayPosts, nil
}

// FetchRandomBeauty randomly fetch a model.Beauty
func FetchRandomBeauty() (model.Beauty, error) {
	prefix := "[正妹]"
	page := rand.Intn(50) + 11 // 10 ~ 60
	posts, err := fetchSearchResult(prefix, page, 90)

	if err != nil {
		return model.Beauty{}, err
	}

	idx := rand.Intn(len(posts)) // 0 ~ len(posts)-1
	p := posts[idx]
	previewImg := fetchPreviewImgURL(p.href)

	b := model.Beauty{
		NVote:      p.nVote,
		Title:      trimTitlePrefix(p.title),
		Href:       p.href,
		PreviewImg: previewImg,
	}
	return b, nil
}

func getBestBeauties(posts []post) []model.Beauty {
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
			imgURL := fetchPreviewImgURL(p.href)
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
		beauties[i].Title = trimTitlePrefix(beauties[i].Title)
	}

	return beauties
}

// FetchBeauties send a request to get beauties from getDailyBeauties api
func FetchBeauties() ([]model.Beauty, error) {
	posts, err := fetchYesterdayPosts()

	if err != nil {
		return nil, err
	}

	beauties := getBestBeauties(posts)

	return beauties, nil
}
