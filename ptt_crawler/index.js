const axios = require('axios')
const cheerio = require('cheerio')
const co = require('co')
const { isToday, isYesterday, isBeforeYesterday } = require('./time')
const { formatNVote } = require('./format')

const genUrl = pageNum => `https://www.ptt.cc/bbs/Beauty/index${pageNum}.html`

// 2558 meas the latest page is https://www.ptt.cc/bbs/Beauty/index2558.html
const getLatestPageNumber = co.wrap(function*() {
  const res = yield axios('https://www.ptt.cc/bbs/Beauty/index.html')
  const body = res.data
  const $ = cheerio.load(body)
  const prevPageHref = $('.btn-group-paging > a:nth-child(2)').attr('href')
  // only get number part in href
  const prevPageNumber = +prevPageHref.match(/\d+/g)[0]
  return prevPageNumber + 1
})

const getPostsInAPage = co.wrap(function*(pageNum) {
  const url = genUrl(pageNum)
  const res = yield axios(url)
  const body = res.data
  const $ = cheerio.load(body)

  const posts = []
  $('div.r-ent').each(function() {
    const nVoteText = $(this)
      .find('.hl')
      .text()
    const nVote = formatNVote(nVoteText)

    posts.push({ nVote })
  })
  return posts
})

const getYesterdayPosts = co.wrap(function*() {
  let n = yield getLatestPageNumber()
  // const url = genUrl(n)
  // const
  // while (true) {
  //   n--
  // }
})

// ref: https://cloud.google.com/functions/docs/writing/http
exports.getDailyBeauties = (_, res) => {
  console.log('----------')
  getPostsInAPage(2556).then(n => {
    res.send({ n })
  })
  // getYesterdayPosts().then(n => {
  //   res.send({ n })
  // })
}
