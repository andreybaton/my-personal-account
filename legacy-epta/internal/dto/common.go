package dto

type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"Operation completed successfully"`
}

// ответ с посещаемостью занятия
type LessonAttendanceResponse struct {
	LessonID   int                       `json:"lesson_id"`
	Discipline string                    `json:"discipline"`
	Date       string                    `json:"date"`
	Attendance []StudentAttendanceDetail `json:"attendance"`
}

type StudentAttendanceDetail struct {
	StudentID       int    `json:"student_id"`
	StudentName     string `json:"student_name"`
	StudentIDNumber string `json:"student_id_number"`
	Status          string `json:"status"`
	RecordedAt      string `json:"recorded_at"`
}

// сводка посещаемости по группе
type GroupAttendanceSummary struct {
	StudentID      int     `json:"student_id"`
	StudentName    string  `json:"student_name"`
	TotalLessons   int     `json:"total_lessons"`
	Attended       int     `json:"attended"`
	AttendanceRate float64 `json:"attendance_rate"`
}
