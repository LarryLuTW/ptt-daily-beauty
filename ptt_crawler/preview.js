const axios = require('axios')
const cheerio = require('cheerio')
const co = require('co')

// const url = 'https://www.ptt.cc/bbs/Beauty/M.1532857638.A.4E0.html'
// imgUrl = https://i.imgur.com/zNcPc5S.jpg
exports.getPreviewImage = co.wrap(function*(url) {
  const { data: body } = yield axios(url)
  const $ = cheerio.load(body)

  // TODO: detect image size and choose best image

  // only get first
  const imgUrl = $('#main-content a[href$=".jpg"],a[href$=".png"]').attr('href')
  // if there's no image, return null
  return imgUrl || null
})

// const url = 'https://www.ptt.cc/bbs/Beauty/M.1532857638.A.4E0.html'
// exports.getPreviewImage(url).then(previewImg => {
//   console.log(previewImg)
// })
