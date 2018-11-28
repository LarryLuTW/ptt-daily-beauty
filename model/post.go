package model

import "time"

type Post struct {
	Title string
	Href  string
	NVote int
	Date  time.Time
}
