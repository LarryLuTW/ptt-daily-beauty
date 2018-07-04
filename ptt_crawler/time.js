const moment = require('moment')

// if current time is 7/04 14:30
// today midnight is 7/04 00:00
// yesterday midnight is 7/03 00:00
// yesterday midnight < today midnight < current time
const currentTime = moment().utcOffset(8)
const todayMidnight = currentTime.clone().startOf('day')
const yesterdayMidnight = todayMidnight.clone().subtract(1, 'days')

// '7/04' -> true
// '7/03' -> false
exports.isToday = text => {
  const m = moment(`${text} +08`, 'M/DD Z')
  return m.isSame(todayMidnight)
}

// '7/04' -> false
// '7/03' -> true
exports.isYesterday = text => {
  const m = moment(`${text} +08`, 'M/DD Z')
  return m.isSame(yesterdayMidnight)
}

// '7/04' -> false
// '7/03' -> false
// '7/02' -> true
exports.isBeforeYesterday = text => {
  const m = moment(`${text} +08`, 'M/DD Z')
  return m.isBefore(yesterdayMidnight)
}
