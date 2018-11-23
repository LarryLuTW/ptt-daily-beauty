package ptt

import "strconv"

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
