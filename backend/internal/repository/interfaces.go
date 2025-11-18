package repository

import (
	"backend/internal/dto"
	"backend/internal/entity"
	"time"
)

type ScheduleRepository interface {
	// методы получения расписания
	GetStudentSchedule(studentID int, startDate, endDate time.Time) ([]entity.Lesson, error)
	GetStudentScheduleWithDetails(studentID int, startDate, endDate time.Time) ([]dto.LessonResponse, error)
	GetGroupSchedule(groupID int, startDate, endDate time.Time) ([]dto.LessonResponse, error)
	GetTeacherSchedule(teacherID int, startDate, endDate time.Time) ([]dto.LessonResponse, error)
	GetClassroomSchedule(classroomID int, startDate, endDate time.Time) ([]dto.LessonResponse, error)

	// методы для работы с занятиями
	GetLessonByID(lessonID int) (*entity.Lesson, error)
	GetLessonWithDetails(lessonID int) (*dto.LessonResponse, error)
	CreateLesson(lesson *entity.Lesson) error
	UpdateLesson(lesson *entity.Lesson) error
	DeleteLesson(lessonID int) error

	// методы проверки конфликтов
	CheckTeacherConflict(teacherID int, date time.Time, startTime, endTime string, excludeLessonID *int) (bool, error)
	CheckClassroomConflict(classroomID int, date time.Time, startTime, endTime string, excludeLessonID *int) (bool, error)
	CheckGroupConflict(groupID int, date time.Time, startTime, endTime string, excludeLessonID *int) (bool, error)

	// доп
	GetLessonsByFilter(filter dto.ScheduleFilter) ([]dto.LessonResponse, error)
	GetCurrentWeekSchedule(studentID int) ([]dto.LessonResponse, error)
	GetTodaySchedule(studentID int) ([]dto.LessonResponse, error)
}

type AttendanceRepository interface {
	// методы посещаемости
	MarkAttendance(attendance *entity.Attendance) error
	MarkAttendanceBatch(attendances []entity.Attendance) error
	GetStudentAttendance(studentID int, startDate, endDate time.Time) ([]entity.Attendance, error)
	GetStudentAttendanceWithDetails(studentID int, startDate, endDate time.Time) ([]dto.AttendanceResponse, error)
	GetLessonAttendance(lessonID int) ([]entity.Attendance, error)
	GetLessonAttendanceWithDetails(lessonID int) ([]struct {
		entity.Attendance
		StudentName     string `db:"student_name"`
		StudentIDNumber string `db:"student_id_number"`
	}, error)
	DeleteAttendance(studentID, lessonID int) error

	// статистика
	GetAttendanceStats(studentID int, startDate, endDate time.Time) (*dto.AttendanceStatsResponse, error)
	GetAttendanceStatsByPeriod(studentID int, period string) (*dto.AttendanceStatsResponse, error)
	GetAttendanceSummaryByGroup(groupID int, startDate, endDate time.Time) ([]struct {
		StudentID      int     `db:"student_id"`
		StudentName    string  `db:"student_name"`
		TotalLessons   int     `db:"total_lessons"`
		Attended       int     `db:"attended"`
		AttendanceRate float64 `db:"attendance_rate"`
	}, error)

	// доп
	GetStudentAttendanceByDiscipline(studentID, disciplineID int, startDate, endDate time.Time) ([]entity.Attendance, error)
	CheckAttendanceExists(studentID, lessonID int) (bool, error)
}

type GroupRepository interface {
	GetGroupByID(groupID int) (*entity.Group, error)
	GetGroupWithDetails(groupID int) (*dto.GroupDetailsResponse, error)
	GetGroupsByFaculty(facultyID int) ([]dto.GroupResponse, error)
	GetAllGroups() ([]dto.GroupResponse, error)
	GetGroupStudents(groupID int) ([]dto.GroupStudentResponse, error)
}

type StudentRepository interface {
	GetStudentByID(studentID int) (*entity.Student, error)
	GetStudentWithDetails(studentID int) (*dto.StudentProfileResponse, error)
	GetStudentsByGroup(groupID int) ([]dto.GroupStudentResponse, error)
}
