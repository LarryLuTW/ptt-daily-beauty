package ptt

import (
	"strconv"
	"strings"
)

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

// [正妹] 大橋未久 -> 大橋未久
func trimTitlePrefix(title string) string {
	return strings.TrimPrefix(title, "[正妹] ")
}
