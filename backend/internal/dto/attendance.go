package dto

type AttendanceResponse struct {
	LessonID   int    `json:"lesson_id"`
	Discipline string `json:"discipline"`
	Teacher    string `json:"teacher"`
	Date       string `json:"date"`
	StartTime  string `json:"start_time"`
	Status     string `json:"status"`
	RecordedAt string `json:"recorded_at"`
}

type AttendanceStatsResponse struct {
	StudentID      int     `json:"student_id"`
	PeriodStart    string  `json:"period_start"`
	PeriodEnd      string  `json:"period_end"`
	TotalLessons   int     `json:"total_lessons"`
	Attended       int     `json:"attended"`
	Absent         int     `json:"absent"`
	Late           int     `json:"late"`
	AttendanceRate float64 `json:"attendance_rate"`
}

type MarkAttendanceRequest struct {
	StudentID int    `json:"student_id" validate:"required"`
	LessonID  int    `json:"lesson_id" validate:"required"`
	Status    string `json:"status" validate:"required,oneof=present absent late"`
}
