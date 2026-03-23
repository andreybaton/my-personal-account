package repository

import (
	"time"

	"backend/internal/dto"
	"backend/internal/entity"

	"github.com/jmoiron/sqlx"
)

type attendanceRepository struct {
	db *sqlx.DB
}

func NewAttendanceRepository(db *sqlx.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}

func (r *attendanceRepository) MarkAttendance(attendance *entity.Attendance) error {
	query := `
        INSERT INTO attendances (student_id, lesson_id, status) 
        VALUES ($1, $2, $3)
        ON CONFLICT (student_id, lesson_id) 
        DO UPDATE SET status = $3, recorded_at = CURRENT_TIMESTAMP
    `

	_, err := r.db.Exec(query, attendance.StudentID, attendance.LessonID, attendance.Status)
	return err
}

func (r *attendanceRepository) MarkAttendanceBatch(attendances []entity.Attendance) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
        INSERT INTO attendances (student_id, lesson_id, status) 
        VALUES ($1, $2, $3)
        ON CONFLICT (student_id, lesson_id) 
        DO UPDATE SET status = $3, recorded_at = CURRENT_TIMESTAMP
    `

	for _, attendance := range attendances {
		_, err := tx.Exec(query, attendance.StudentID, attendance.LessonID, attendance.Status)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *attendanceRepository) GetStudentAttendance(studentID int, startDate, endDate time.Time) ([]entity.Attendance, error) {
	query := `
        SELECT a.* 
        FROM attendances a
        JOIN lessons l ON a.lesson_id = l.id
        WHERE a.student_id = $1 
        AND l.lesson_date BETWEEN $2 AND $3
        ORDER BY l.lesson_date, l.start_time
    `

	var attendances []entity.Attendance
	err := r.db.Select(&attendances, query, studentID, startDate, endDate)
	return attendances, err
}

func (r *attendanceRepository) GetStudentAttendanceWithDetails(studentID int, startDate, endDate time.Time) ([]dto.AttendanceResponse, error) {
	query := `
        SELECT 
            a.lesson_id,
            d.name as discipline,
            u.name as teacher,
            l.lesson_date as date,
            l.start_time,
            a.status,
            a.recorded_at
        FROM attendances a
        JOIN lessons l ON a.lesson_id = l.id
        JOIN disciplines d ON l.discipline_id = d.id
        JOIN teachers t ON l.teacher_id = t.id
        JOIN employees e ON t.employ_id = e.id
        JOIN users u ON e.user_id = u.id
        WHERE a.student_id = $1 
        AND l.lesson_date BETWEEN $2 AND $3
        ORDER BY l.lesson_date, l.start_time
    `

	var attendanceDetails []dto.AttendanceResponse
	err := r.db.Select(&attendanceDetails, query, studentID, startDate, endDate)
	return attendanceDetails, err
}

func (r *attendanceRepository) GetLessonAttendance(lessonID int) ([]entity.Attendance, error) {
	query := `
        SELECT a.* 
        FROM attendances a
        WHERE a.lesson_id = $1
        ORDER BY a.student_id
    `

	var attendances []entity.Attendance
	err := r.db.Select(&attendances, query, lessonID)
	return attendances, err
}

func (r *attendanceRepository) GetLessonAttendanceWithDetails(lessonID int) ([]struct {
	entity.Attendance
	StudentName     string `db:"student_name"`
	StudentIDNumber string `db:"student_id_number"`
}, error) {
	query := `
        SELECT 
            a.*,
            u.name as student_name,
            s.student_id_number
        FROM attendances a
        JOIN students s ON a.student_id = s.id
        JOIN users u ON s.user_id = u.id
        WHERE a.lesson_id = $1
        ORDER BY u.name
    `

	var result []struct {
		entity.Attendance
		StudentName     string `db:"student_name"`
		StudentIDNumber string `db:"student_id_number"`
	}

	err := r.db.Select(&result, query, lessonID)
	return result, err
}

func (r *attendanceRepository) DeleteAttendance(studentID, lessonID int) error {
	query := `DELETE FROM attendances WHERE student_id = $1 AND lesson_id = $2`

	_, err := r.db.Exec(query, studentID, lessonID)
	return err
}

func (r *attendanceRepository) GetAttendanceStats(studentID int, startDate, endDate time.Time) (*dto.AttendanceStatsResponse, error) {
	query := `
        SELECT 
            COUNT(*) as total_lessons,
            COUNT(CASE WHEN a.status = 'present' THEN 1 END) as attended,
            COUNT(CASE WHEN a.status = 'absent' THEN 1 END) as absent,
            COUNT(CASE WHEN a.status = 'late' THEN 1 END) as late
        FROM lessons l
        JOIN groups g ON l.group_id = g.id
        JOIN students s ON s.group_id = g.id
        LEFT JOIN attendances a ON l.id = a.lesson_id AND a.student_id = $1
        WHERE s.id = $1 AND l.lesson_date BETWEEN $2 AND $3
    `

	var stats struct {
		TotalLessons int `db:"total_lessons"`
		Attended     int `db:"attended"`
		Absent       int `db:"absent"`
		Late         int `db:"late"`
	}

	err := r.db.Get(&stats, query, studentID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	attendanceRate := 0.0
	if stats.TotalLessons > 0 {
		attendanceRate = float64(stats.Attended) / float64(stats.TotalLessons) * 100
	}

	return &dto.AttendanceStatsResponse{
		StudentID:      studentID,
		PeriodStart:    startDate.Format("2006-01-02"),
		PeriodEnd:      endDate.Format("2006-01-02"),
		TotalLessons:   stats.TotalLessons,
		Attended:       stats.Attended,
		Absent:         stats.Absent,
		Late:           stats.Late,
		AttendanceRate: attendanceRate,
	}, nil
}

func (r *attendanceRepository) GetAttendanceStatsByPeriod(studentID int, period string) (*dto.AttendanceStatsResponse, error) {
	var startDate, endDate time.Time
	now := time.Now()

	switch period {
	case "week":
		startDate = getStartOfWeek(now)
		endDate = startDate.AddDate(0, 0, 6)
	case "month":
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endDate = startDate.AddDate(0, 1, -1)
	case "semester":
		if now.Month() >= time.September || now.Month() <= time.January {
			startDate = time.Date(now.Year(), time.September, 1, 0, 0, 0, 0, now.Location())
			endDate = time.Date(now.Year()+1, time.January, 31, 0, 0, 0, 0, now.Location())
		} else {
			startDate = time.Date(now.Year(), time.February, 1, 0, 0, 0, 0, now.Location())
			endDate = time.Date(now.Year(), time.June, 30, 0, 0, 0, 0, now.Location())
		}
	default:
		startDate = getStartOfWeek(now)
		endDate = startDate.AddDate(0, 0, 6)
	}

	return r.GetAttendanceStats(studentID, startDate, endDate)
}

func (r *attendanceRepository) GetStudentAttendanceByDiscipline(studentID, disciplineID int, startDate, endDate time.Time) ([]entity.Attendance, error) {
	query := `
        SELECT a.* 
        FROM attendances a
        JOIN lessons l ON a.lesson_id = l.id
        WHERE a.student_id = $1 
        AND l.discipline_id = $2
        AND l.lesson_date BETWEEN $3 AND $4
        ORDER BY l.lesson_date, l.start_time
    `

	var attendances []entity.Attendance
	err := r.db.Select(&attendances, query, studentID, disciplineID, startDate, endDate)
	return attendances, err
}

func (r *attendanceRepository) CheckAttendanceExists(studentID, lessonID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM attendances WHERE student_id = $1 AND lesson_id = $2)`

	var exists bool
	err := r.db.Get(&exists, query, studentID, lessonID)
	return exists, err
}

func (r *attendanceRepository) GetAttendanceSummaryByGroup(groupID int, startDate, endDate time.Time) ([]struct {
	StudentID      int     `db:"student_id"`
	StudentName    string  `db:"student_name"`
	TotalLessons   int     `db:"total_lessons"`
	Attended       int     `db:"attended"`
	AttendanceRate float64 `db:"attendance_rate"`
}, error) {
	query := `
        SELECT 
            s.id as student_id,
            u.name as student_name,
            COUNT(l.id) as total_lessons,
            COUNT(CASE WHEN a.status = 'present' THEN 1 END) as attended,
            CASE 
                WHEN COUNT(l.id) > 0 THEN 
                    ROUND(COUNT(CASE WHEN a.status = 'present' THEN 1 END) * 100.0 / COUNT(l.id), 2)
                ELSE 0 
            END as attendance_rate
        FROM students s
        JOIN users u ON s.user_id = u.id
        JOIN lessons l ON l.group_id = s.group_id
        LEFT JOIN attendances a ON l.id = a.lesson_id AND a.student_id = s.id
        WHERE s.group_id = $1 
        AND l.lesson_date BETWEEN $2 AND $3
        GROUP BY s.id, u.name
        ORDER BY attendance_rate DESC, u.name
    `

	var summary []struct {
		StudentID      int     `db:"student_id"`
		StudentName    string  `db:"student_name"`
		TotalLessons   int     `db:"total_lessons"`
		Attended       int     `db:"attended"`
		AttendanceRate float64 `db:"attendance_rate"`
	}

	err := r.db.Select(&summary, query, groupID, startDate, endDate)
	return summary, err
}
