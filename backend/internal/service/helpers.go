package service

import "time"

// русское название дня недели
func getRussianDayName(day time.Weekday) string {
	days := map[time.Weekday]string{
		time.Monday:    "Понедельник",
		time.Tuesday:   "Вторник",
		time.Wednesday: "Среда",
		time.Thursday:  "Четверг",
		time.Friday:    "Пятница",
		time.Saturday:  "Суббота",
		time.Sunday:    "Воскресенье",
	}
	return days[day]
}

func getStartOfWeek(date time.Time) time.Time {
	offset := (int(date.Weekday()) - 1) % 7
	if offset < 0 {
		offset += 7
	}
	return date.AddDate(0, 0, -offset)
}
