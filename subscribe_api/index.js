// ref: https://cloud.google.com/functions/docs/writing/http
exports.subscribe = co.wrap(function*(req, res) {
  res.send('subscribe success')
})

exports.unsubscribe = co.wrap(function*(req, res) {
  res.send('unsubscribe success')
})
