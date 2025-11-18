package service

import (
	"backend/internal/dto"
)

type ScheduleService interface {
	GetWeeklySchedule(studentID int, weekStart string) (*dto.WeeklyScheduleResponse, error)
	GetDailySchedule(studentID int, date string) (*dto.DayScheduleResponse, error)
	GetStudentSchedule(studentID int, startDate, endDate string) ([]dto.LessonResponse, error)
	GetGroupSchedule(groupID int, startDate, endDate string) ([]dto.LessonResponse, error)
	GetTeacherSchedule(teacherID int, startDate, endDate string) ([]dto.LessonResponse, error)
	GetClassroomSchedule(classroomID int, startDate, endDate string) ([]dto.LessonResponse, error)

	CreateLesson(request dto.CreateLessonRequest) error
	UpdateLesson(lessonID int, request dto.CreateLessonRequest) error
	DeleteLesson(lessonID int) error

	ValidateLessonConflict(request dto.CreateLessonRequest) error
	GetLessonDetails(lessonID int) (*dto.LessonResponse, error)
}

type AttendanceService interface {
	MarkAttendance(request dto.MarkAttendanceRequest) error
	MarkAttendanceBatch(requests []dto.MarkAttendanceRequest) error
	GetStudentAttendance(studentID int, startDate, endDate string) ([]dto.AttendanceResponse, error)
	GetAttendanceStats(studentID int, period string) (*dto.AttendanceStatsResponse, error)
	GetLessonAttendance(lessonID int) (interface{}, error)
	GetGroupAttendanceSummary(groupID int, startDate, endDate string) (interface{}, error)
}

type GroupService interface {
	GetGroupDetails(groupID int) (*dto.GroupDetailsResponse, error)
	GetGroupsByFaculty(facultyID int) ([]dto.GroupResponse, error)
	GetAllGroups() ([]dto.GroupResponse, error)
}

type StudentService interface {
	GetStudentProfile(studentID int) (*dto.StudentProfileResponse, error)
	GetStudentAcademicInfo(studentID int) (*dto.StudentAcademicResponse, error)
}
