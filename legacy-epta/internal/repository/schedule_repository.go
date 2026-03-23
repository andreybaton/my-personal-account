package repository

import (
	"database/sql"
	"fmt"
	"time"

	"backend/internal/dto"
	"backend/internal/entity"

	"github.com/jmoiron/sqlx"
)

type scheduleRepository struct {
	db *sqlx.DB
}

func NewScheduleRepository(db *sqlx.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) GetStudentSchedule(studentID int, startDate, endDate time.Time) ([]entity.Lesson, error) {
	query := `
        SELECT l.* 
        FROM lessons l
        JOIN groups g ON l.group_id = g.id
        JOIN students s ON s.group_id = g.id
        WHERE s.id = $1 AND l.lesson_date BETWEEN $2 AND $3
        ORDER BY l.lesson_date, l.start_time
    `

	var lessons []entity.Lesson
	err := r.db.Select(&lessons, query, studentID, startDate, endDate)
	return lessons, err
}

func (r *scheduleRepository) GetStudentScheduleWithDetails(studentID int, startDate, endDate time.Time) ([]dto.LessonResponse, error) {
	query := `
        SELECT 
            l.id,
            l.lesson_date as date,
            l.start_time,
            l.end_time,
            l.type,
            d.name as discipline,
            u.name as teacher,
            t.academic_degree as teacher_degree,
            c.room_number as classroom,
            b.type as building,
            b.address
        FROM lessons l
        JOIN disciplines d ON l.discipline_id = d.id
        JOIN teachers t ON l.teacher_id = t.id
        JOIN employees e ON t.employ_id = e.id
        JOIN users u ON e.user_id = u.id
        JOIN classrooms c ON l.classroom_id = c.id
        JOIN buildings b ON c.building_id = b.id
        JOIN groups g ON l.group_id = g.id
        JOIN students s ON s.group_id = g.id
        WHERE s.id = $1 AND l.lesson_date BETWEEN $2 AND $3
        ORDER BY l.lesson_date, l.start_time
    `

	var lessons []dto.LessonResponse
	err := r.db.Select(&lessons, query, studentID, startDate, endDate)
	return lessons, err
}

func (r *scheduleRepository) GetGroupSchedule(groupID int, startDate, endDate time.Time) ([]dto.LessonResponse, error) {
	query := `
        SELECT 
            l.id,
            l.lesson_date as date,
            l.start_time,
            l.end_time,
            l.type,
            d.name as discipline,
            u.name as teacher,
            t.academic_degree as teacher_degree,
            c.room_number as classroom,
            b.type as building,
            b.address
        FROM lessons l
        JOIN disciplines d ON l.discipline_id = d.id
        JOIN teachers t ON l.teacher_id = t.id
        JOIN employees e ON t.employ_id = e.id
        JOIN users u ON e.user_id = u.id
        JOIN classrooms c ON l.classroom_id = c.id
        JOIN buildings b ON c.building_id = b.id
        WHERE l.group_id = $1 AND l.lesson_date BETWEEN $2 AND $3
        ORDER BY l.lesson_date, l.start_time
    `

	var lessons []dto.LessonResponse
	err := r.db.Select(&lessons, query, groupID, startDate, endDate)
	return lessons, err
}

func (r *scheduleRepository) GetTeacherSchedule(teacherID int, startDate, endDate time.Time) ([]dto.LessonResponse, error) {
	query := `
        SELECT 
            l.id,
            l.lesson_date as date,
            l.start_time,
            l.end_time,
            l.type,
            d.name as discipline,
            u.name as teacher,
            t.academic_degree as teacher_degree,
            c.room_number as classroom,
            b.type as building,
            b.address,
            g.name as group_name
        FROM lessons l
        JOIN disciplines d ON l.discipline_id = d.id
        JOIN teachers t ON l.teacher_id = t.id
        JOIN employees e ON t.employ_id = e.id
        JOIN users u ON e.user_id = u.id
        JOIN classrooms c ON l.classroom_id = c.id
        JOIN buildings b ON c.building_id = b.id
        JOIN groups g ON l.group_id = g.id
        WHERE l.teacher_id = $1 AND l.lesson_date BETWEEN $2 AND $3
        ORDER BY l.lesson_date, l.start_time
    `

	var lessons []dto.LessonResponse
	err := r.db.Select(&lessons, query, teacherID, startDate, endDate)
	return lessons, err
}

func (r *scheduleRepository) GetClassroomSchedule(classroomID int, startDate, endDate time.Time) ([]dto.LessonResponse, error) {
	query := `
        SELECT 
            l.id,
            l.lesson_date as date,
            l.start_time,
            l.end_time,
            l.type,
            d.name as discipline,
            u.name as teacher,
            t.academic_degree as teacher_degree,
            c.room_number as classroom,
            b.type as building,
            b.address,
            g.name as group_name
        FROM lessons l
        JOIN disciplines d ON l.discipline_id = d.id
        JOIN teachers t ON l.teacher_id = t.id
        JOIN employees e ON t.employ_id = e.id
        JOIN users u ON e.user_id = u.id
        JOIN classrooms c ON l.classroom_id = c.id
        JOIN buildings b ON c.building_id = b.id
        JOIN groups g ON l.group_id = g.id
        WHERE l.classroom_id = $1 AND l.lesson_date BETWEEN $2 AND $3
        ORDER BY l.lesson_date, l.start_time
    `

	var lessons []dto.LessonResponse
	err := r.db.Select(&lessons, query, classroomID, startDate, endDate)
	return lessons, err
}

func (r *scheduleRepository) GetLessonByID(lessonID int) (*entity.Lesson, error) {
	query := `SELECT * FROM lessons WHERE id = $1`

	var lesson entity.Lesson
	err := r.db.Get(&lesson, query, lessonID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &lesson, err
}

func (r *scheduleRepository) GetLessonWithDetails(lessonID int) (*dto.LessonResponse, error) {
	query := `
        SELECT 
            l.id,
            l.lesson_date as date,
            l.start_time,
            l.end_time,
            l.type,
            d.name as discipline,
            u.name as teacher,
            t.academic_degree as teacher_degree,
            c.room_number as classroom,
            b.type as building,
            b.address,
            g.name as group_name
        FROM lessons l
        JOIN disciplines d ON l.discipline_id = d.id
        JOIN teachers t ON l.teacher_id = t.id
        JOIN employees e ON t.employ_id = e.id
        JOIN users u ON e.user_id = u.id
        JOIN classrooms c ON l.classroom_id = c.id
        JOIN buildings b ON c.building_id = b.id
        JOIN groups g ON l.group_id = g.id
        WHERE l.id = $1
    `

	var lesson dto.LessonResponse
	err := r.db.Get(&lesson, query, lessonID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &lesson, err
}

func (r *scheduleRepository) CreateLesson(lesson *entity.Lesson) error {
	query := `
        INSERT INTO lessons (teacher_id, classroom_id, discipline_id, group_id, lesson_date, start_time, end_time, type)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id
    `

	return r.db.QueryRow(
		query,
		lesson.TeacherID,
		lesson.ClassroomID,
		lesson.DisciplineID,
		lesson.GroupID,
		lesson.LessonDate,
		lesson.StartTime,
		lesson.EndTime,
		lesson.Type,
	).Scan(&lesson.ID)
}

func (r *scheduleRepository) UpdateLesson(lesson *entity.Lesson) error {
	query := `
        UPDATE lessons 
        SET teacher_id = $1, classroom_id = $2, discipline_id = $3, group_id = $4, 
            lesson_date = $5, start_time = $6, end_time = $7, type = $8
        WHERE id = $9
    `

	_, err := r.db.Exec(
		query,
		lesson.TeacherID,
		lesson.ClassroomID,
		lesson.DisciplineID,
		lesson.GroupID,
		lesson.LessonDate,
		lesson.StartTime,
		lesson.EndTime,
		lesson.Type,
		lesson.ID,
	)
	return err
}

func (r *scheduleRepository) DeleteLesson(lessonID int) error {
	query := `DELETE FROM lessons WHERE id = $1`
	_, err := r.db.Exec(query, lessonID)
	return err
}

func (r *scheduleRepository) CheckTeacherConflict(teacherID int, date time.Time, startTime, endTime string, excludeLessonID *int) (bool, error) {
	query := `
        SELECT EXISTS(
            SELECT 1 FROM lessons 
            WHERE teacher_id = $1 
            AND lesson_date = $2 
            AND (
                (start_time <= $3 AND end_time > $3) OR
                (start_time < $4 AND end_time >= $4) OR
                (start_time >= $3 AND end_time <= $4)
            )
            AND ($5 IS NULL OR id != $5)
        )
    `

	var exists bool
	err := r.db.Get(&exists, query, teacherID, date, startTime, endTime, excludeLessonID)
	return exists, err
}

func (r *scheduleRepository) CheckClassroomConflict(classroomID int, date time.Time, startTime, endTime string, excludeLessonID *int) (bool, error) {
	query := `
        SELECT EXISTS(
            SELECT 1 FROM lessons 
            WHERE classroom_id = $1 
            AND lesson_date = $2 
            AND (
                (start_time <= $3 AND end_time > $3) OR
                (start_time < $4 AND end_time >= $4) OR
                (start_time >= $3 AND end_time <= $4)
            )
            AND ($5 IS NULL OR id != $5)
        )
    `

	var exists bool
	err := r.db.Get(&exists, query, classroomID, date, startTime, endTime, excludeLessonID)
	return exists, err
}

func (r *scheduleRepository) CheckGroupConflict(groupID int, date time.Time, startTime, endTime string, excludeLessonID *int) (bool, error) {
	query := `
        SELECT EXISTS(
            SELECT 1 FROM lessons 
            WHERE group_id = $1 
            AND lesson_date = $2 
            AND (
                (start_time <= $3 AND end_time > $3) OR
                (start_time < $4 AND end_time >= $4) OR
                (start_time >= $3 AND end_time <= $4)
            )
            AND ($5 IS NULL OR id != $5)
        )
    `

	var exists bool
	err := r.db.Get(&exists, query, groupID, date, startTime, endTime, excludeLessonID)
	return exists, err
}

func (r *scheduleRepository) GetLessonsByFilter(filter dto.ScheduleFilter) ([]dto.LessonResponse, error) {
	query := `
        SELECT 
            l.id,
            l.lesson_date as date,
            l.start_time,
            l.end_time,
            l.type,
            d.name as discipline,
            u.name as teacher,
            t.academic_degree as teacher_degree,
            c.room_number as classroom,
            b.type as building,
            b.address,
            g.name as group_name
        FROM lessons l
        JOIN disciplines d ON l.discipline_id = d.id
        JOIN teachers t ON l.teacher_id = t.id
        JOIN employees e ON t.employ_id = e.id
        JOIN users u ON e.user_id = u.id
        JOIN classrooms c ON l.classroom_id = c.id
        JOIN buildings b ON c.building_id = b.id
        JOIN groups g ON l.group_id = g.id
        WHERE 1=1
    `

	args := []interface{}{}
	argIndex := 1

	if filter.StartDate != "" {
		query += fmt.Sprintf(" AND l.lesson_date >= $%d", argIndex)
		args = append(args, filter.StartDate)
		argIndex++
	}

	if filter.EndDate != "" {
		query += fmt.Sprintf(" AND l.lesson_date <= $%d", argIndex)
		args = append(args, filter.EndDate)
		argIndex++
	}

	if filter.GroupID != 0 {
		query += fmt.Sprintf(" AND l.group_id = $%d", argIndex)
		args = append(args, filter.GroupID)
		argIndex++
	}

	if filter.TeacherID != 0 {
		query += fmt.Sprintf(" AND l.teacher_id = $%d", argIndex)
		args = append(args, filter.TeacherID)
		argIndex++
	}

	if filter.BuildingID != 0 {
		query += fmt.Sprintf(" AND c.building_id = $%d", argIndex)
		args = append(args, filter.BuildingID)
		argIndex++
	}

	query += " ORDER BY l.lesson_date, l.start_time"

	var lessons []dto.LessonResponse
	err := r.db.Select(&lessons, query, args...)
	return lessons, err
}

func (r *scheduleRepository) GetCurrentWeekSchedule(studentID int) ([]dto.LessonResponse, error) {
	now := time.Now()
	startOfWeek := getStartOfWeek(now)
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	return r.GetStudentScheduleWithDetails(studentID, startOfWeek, endOfWeek)
}

func (r *scheduleRepository) GetTodaySchedule(studentID int) ([]dto.LessonResponse, error) {
	today := time.Now().Format("2006-01-02")
	startDate, _ := time.Parse("2006-01-02", today)
	endDate := startDate

	return r.GetStudentScheduleWithDetails(studentID, startDate, endDate)
}
