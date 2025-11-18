package entity

import "time"

type Classroom struct {
	ID         int    `db:"id"`
	Type       string `db:"type"`
	BuildingID int    `db:"building_id"`
	RoomNumber string `db:"room_number"`
}

type Discipline struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	DepartmentID int    `db:"department_id"`
}

type TeacherDiscipline struct {
	TeacherID    int `db:"teacher_id"`
	DisciplineID int `db:"discipline_id"`
}

type Lesson struct {
	ID           int       `db:"id"`
	TeacherID    int       `db:"teacher_id"`
	ClassroomID  int       `db:"classroom_id"`
	DisciplineID int       `db:"discipline_id"`
	GroupID      int       `db:"group_id"`
	LessonDate   time.Time `db:"lesson_date"`
	StartTime    string    `db:"start_time"`
	EndTime      string    `db:"end_time"`
	Type         string    `db:"type"`
}

type Attendance struct {
	StudentID  int       `db:"student_id"`
	LessonID   int       `db:"lesson_id"`
	Status     string    `db:"status"`
	RecordedAt time.Time `db:"recorded_at"`
}

type Grade struct {
	ID           int       `db:"id"`
	StudentID    int       `db:"student_id"`
	DisciplineID int       `db:"discipline_id"`
	LessonID     *int      `db:"lesson_id"`   // может быть NULL
	GradeValue   *float64  `db:"grade_value"` // может быть NULL
	GradeType    string    `db:"grade_type"`
	Semester     string    `db:"semester"`
	CreatedAt    time.Time `db:"created_at"`
}

type Curriculum struct {
	ID               int  `db:"id"`
	SpecializationID int  `db:"specialization_id"`
	DisciplineID     int  `db:"discipline_id"`
	SemesterNumber   int  `db:"semester_number"`
	HoursTotal       int  `db:"hours_total"`
	HoursLecture     *int `db:"hours_lecture"`  // может быть NULL
	HoursPractice    *int `db:"hours_practice"` // может быть NULL
	IsRequired       bool `db:"is_required"`
}
