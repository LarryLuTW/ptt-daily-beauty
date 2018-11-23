package ptt

import "time"

func isToday(t time.Time) bool {
	loc, _ := time.LoadLocation("Asia/Taipei")
	current := time.Now().In(loc)
	return t.YearDay() == current.YearDay()
}

func isYesterday(t time.Time) bool {
	// FIXME: 跨年 1/1 < 12/31
	loc, _ := time.LoadLocation("Asia/Taipei")
	current := time.Now().In(loc)
	return t.YearDay() == current.YearDay()-1
}

func isBeforeYesterday(t time.Time) bool {
	// FIXME: 跨年 1/1 < 12/31
	loc, _ := time.LoadLocation("Asia/Taipei")
	current := time.Now().In(loc)
	return t.YearDay() < current.YearDay()-1
}
