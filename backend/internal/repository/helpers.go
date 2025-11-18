package repository

import "time"

func getStartOfWeek(date time.Time) time.Time {
	offset := (int(date.Weekday()) - 1) % 7
	if offset < 0 {
		offset += 7
	}
	return date.AddDate(0, 0, -offset)
}
