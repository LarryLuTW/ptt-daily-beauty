package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"main/model"
)

const url = "https://us-central1-daily-beauty-209105.cloudfunctions.net/getDailyBeauties"

// FetchBeauties send a request to get beauties from getDailyBeauties api
func FetchBeauties() ([]model.Beauty, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	beauties := make([]model.Beauty, 3)
	err = json.Unmarshal(body, &beauties)

	if err != nil {
		return nil, err
	}

	beauties[0].Rank = "一"
	beauties[1].Rank = "二"
	beauties[2].Rank = "三"

	// [正妹] 大橋未久 -> 大橋未久
	for i := range beauties {
		beauties[i].Title = beauties[i].Title[9:]
	}

	return beauties, nil
}
