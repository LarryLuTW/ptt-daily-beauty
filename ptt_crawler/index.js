const axios = require('axios')
const cheerio = require('cheerio')
const co = require('co')
const { isYesterday, isBeforeYesterday } = require('./time')
const { formatNVote } = require('./format')
const { getPreviewImage } = require('./preview')

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
    const nVoteEl = $(this).find('.hl')
    const nVote = formatNVote(nVoteEl.text())

    const titleEl = $(this).find('.title > a')
    const title = titleEl.text()
    const href = 'https://www.ptt.cc' + titleEl.attr('href')

    const date = $(this)
      .find('.meta .date')
      .text()
      .trim()

    const mark = $(this)
      .find('.mark')
      .text()

    posts.push({ nVote, title, href, mark, date })
  })
  return posts
})

const getYesterdayPosts = co.wrap(function*() {
  let n = yield getLatestPageNumber()
  const allPosts = []
  while (true) {
    const posts = yield getPostsInAPage(n)
    posts.reverse()
    // posts[0] is latest in a page
    // posts[len-1] is oldest in a page
    allPosts.push(...posts)

    if (posts[0].mark === '' && isBeforeYesterday(posts[0].date)) {
      // if date is the day before yesterday, stop crawling
      break
    }

    n--
  }
  // from old to latest
  allPosts.reverse()
  return allPosts
})

const filter = posts => {
  return (
    posts
      // filter out 公告
      .filter(p => p.mark === '')
      .filter(p => isYesterday(p.date))
      // only find beauty
      .filter(p => p.title.indexOf('[正妹]') !== -1)
  )
}

const sort = posts => posts.sort((a, b) => b.nVote - a.nVote)

// ref: https://cloud.google.com/functions/docs/writing/http
exports.getDailyBeauties = co.wrap(function*(_, res) {
  console.log('----------')
  const allPosts = yield getYesterdayPosts()
  const filteredPosts = filter(allPosts)
  const sortedPosts = sort(filteredPosts)
  const champions = sortedPosts.slice(0, 3)
  // champions = [
  //   {
  //     nVote: 54,
  //     title: '[正妹] 實況主',
  //     href: 'https://www.ptt.cc/bbs/Beauty/M.1532857638.A.4E0.html',
  //     mark: '',
  //     date: '7/29',
  //     previewImg: 'https://i.imgur.com/7YPKLND.jpg'
  //   },
  //   {...},
  //   {...},
  // ]

  // add previewImg property
  for (let i = 0; i < champions.length; i++) {
    champions[i].previewImg = yield getPreviewImage(champions[i].href)
  }

  console.log(JSON.stringify(champions, null, 4))
  res.send(champions)
})
