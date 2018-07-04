// '' -> 0
// 'X3' -> -3
// '爆' -> 100
// '50' -> 50
exports.formatNVote = nVoteText => {
  if (nVoteText.length === 0) {
    return 0
  }
  if (nVoteText[0] === 'X') {
    return -nVoteText.slice(1, nVoteText.length)
  }
  if (nVoteText[0] === '爆') {
    return 100
  }
  return +nVoteText
}
