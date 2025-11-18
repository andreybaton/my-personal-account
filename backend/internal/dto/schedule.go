package dto

type ScheduleRequest struct {
	StartDate string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate   string `json:"end_date" validate:"required,datetime=2006-01-02"`
}

type LessonResponse struct {
	ID            int    `json:"id"`
	Discipline    string `json:"discipline"`
	Teacher       string `json:"teacher"`
	TeacherDegree string `json:"teacher_degree,omitempty"`
	Classroom     string `json:"classroom"`
	Building      string `json:"building"`
	Address       string `json:"address,omitempty"`
	Date          string `json:"date"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	Type          string `json:"type"`
}

type WeeklyScheduleResponse struct {
	WeekNumber int                   `json:"week_number"`
	StartDate  string                `json:"start_date"`
	EndDate    string                `json:"end_date"`
	Days       []DayScheduleResponse `json:"days"`
}

type DayScheduleResponse struct {
	Date    string           `json:"date"`
	DayName string           `json:"day_name"`
	Lessons []LessonResponse `json:"lessons"`
}

type ScheduleFilter struct {
	StartDate  string `form:"start_date"`
	EndDate    string `form:"end_date"`
	Date       string `form:"date"`
	Week       string `form:"week"`
	GroupID    int    `form:"group_id"`
	TeacherID  int    `form:"teacher_id"`
	BuildingID int    `form:"building_id"`
}

type CreateLessonRequest struct {
	TeacherID    int    `json:"teacher_id" validate:"required"`
	ClassroomID  int    `json:"classroom_id" validate:"required"`
	DisciplineID int    `json:"discipline_id" validate:"required"`
	GroupIDs     []int  `json:"group_ids" validate:"required,min=1"` // для нескольких групп
	LessonDate   string `json:"lesson_date" validate:"required,datetime=2006-01-02"`
	StartTime    string `json:"start_time" validate:"required"`
	EndTime      string `json:"end_time" validate:"required"`
	Type         string `json:"type" validate:"required,oneof=lecture practice lab seminar"`
}

type TeacherResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	AcademicDegree string `json:"academic_degree"`
	Department     string `json:"department"`
	Email          string `json:"email"`
}
