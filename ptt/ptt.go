package ptt

import (
	"main/model"
	"main/ptt/api"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

// TODO: split ptt api layer and utils layer
func init() {
	rand.Seed(time.Now().UnixNano())
}

// [正妹] 大橋未久 -> 大橋未久
func trimTitlePrefix(title string) string {
	return strings.TrimPrefix(title, "[正妹] ")
}

func fetchYesterdayPosts() ([]model.Post, error) {
	prefix := "[正妹]"
	recentPosts := make([]model.Post, 0, 20)

	// get recent posts
	page, err := api.FetchPageAmount()
	if err != nil {
		return nil, err
	}

	for ; ; page-- {
		posts, err := api.FetchPage(prefix, page)

		if err != nil {
			return nil, err
		}

		recentPosts = append(recentPosts, posts...)
		oldestDate := recentPosts[len(recentPosts)-1].Date
		if isBeforeYesterday(oldestDate) {
			break
		}
	}

	// filter yesterday post
	yesterdayPosts := make([]model.Post, 0, 10)
	for _, p := range recentPosts {
		if isYesterday(p.Date) {
			yesterdayPosts = append(yesterdayPosts, p)
		}
	}

	return yesterdayPosts, nil
}

// FetchRandomBeauty randomly fetch a model.Beauty
func FetchRandomBeauty() (model.Beauty, error) {
	prefix := "[正妹]"
	page := rand.Intn(50) + 11 // 10 ~ 60
	posts, err := api.Search(prefix, page, 90)

	if err != nil {
		return model.Beauty{}, err
	}

	idx := rand.Intn(len(posts)) // 0 ~ len(posts)-1
	p := posts[idx]
	previewImg := api.FetchPreviewImg(p)

	b := model.Beauty{
		NVote:      p.NVote,
		Title:      trimTitlePrefix(p.Title),
		Href:       p.Href,
		PreviewImg: previewImg,
	}
	return b, nil
}

func getBestBeauties(posts []model.Post) []model.Beauty {
	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].NVote > posts[j].NVote
	})

	champions := posts[:3]
	beauties := make([]model.Beauty, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	for i, p := range champions {
		go func(i int, p model.Post) {
			defer wg.Done()
			imgURL := api.FetchPreviewImg(p)
			beauties[i] = model.Beauty{
				NVote:      p.NVote,
				Title:      p.Title,
				Href:       p.Href,
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
