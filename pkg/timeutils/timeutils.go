package timeutils

import "time"

func GetTodayStart() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func GetMonthStart() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}
