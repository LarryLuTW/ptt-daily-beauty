package model

// Beauty is a struct from getDailyBeauties api
type Beauty struct {
	Rank       string // "一", "二", "三"
	NVote      int
	Title      string
	Href       string
	Date       string
	PreviewImg string
}
