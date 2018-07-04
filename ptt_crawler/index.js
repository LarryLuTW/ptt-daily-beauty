const axios = require('axios')
const cheerio = require('cheerio')

const genUrl = page => `https://www.ptt.cc/bbs/Beauty/index${page}.html`

// 2558 meas the latest page is https://www.ptt.cc/bbs/Beauty/index2558.html
const getLatestPageNumber = () =>
  axios('https://www.ptt.cc/bbs/Beauty/index.html')
    .then(res => res.data)
    .then(body => {
      const $ = cheerio.load(body)
      const prevPageHref = $('.btn-group-paging > a:nth-child(2)').attr('href')
      // only get number part in href
      const prevPageNumber = +prevPageHref.match(/\d+/g)[0]
      return prevPageNumber + 1
    })

// ref: https://cloud.google.com/functions/docs/writing/http
exports.getDailyBeauties = (_, res) => {
  console.log('----------')
  getLatestPageNumber().then(n => {
    res.send({ n })
  })
}
