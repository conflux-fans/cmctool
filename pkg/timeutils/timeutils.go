package timeutils

import "time"

func GetTodayStart(loc *time.Location) time.Time {
	now := time.Now().In(loc)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}

func GetMonthStart(loc *time.Location) time.Time {
	now := time.Now().In(loc)
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc)
}

func ToChinaTime(t time.Time) time.Time {
	return t.In(GetChinaLocation())
}

func GetChinaLocation() *time.Location {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	return loc
}
