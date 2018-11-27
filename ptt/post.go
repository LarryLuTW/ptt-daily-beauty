package ptt

import "time"

type post struct {
	title string
	href  string
	nVote int
	date  time.Time
}
