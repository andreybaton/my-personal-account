package repository

import (
	"database/sql"

	"backend/internal/dto"
	"backend/internal/entity"

	"github.com/jmoiron/sqlx"
)

type studentRepository struct {
	db *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) GetStudentByID(studentID int) (*entity.Student, error) {
	query := `SELECT * FROM students WHERE id = $1`

	var student entity.Student
	err := r.db.Get(&student, query, studentID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &student, err
}

func (r *studentRepository) GetStudentWithDetails(studentID int) (*dto.StudentProfileResponse, error) {
	query := `
        SELECT 
            s.id,
            s.user_id,
            s.student_id_number,
            u.name,
            u.email,
            s.group_id,
            g.name as group_name,
            f.id as faculty_id,
            f.name as faculty_name,
            g.specialization,
            g.admission_year
        FROM students s
        JOIN users u ON s.user_id = u.id
        LEFT JOIN groups g ON s.group_id = g.id
        LEFT JOIN faculties f ON g.faculty_id = f.id
        WHERE s.id = $1
    `

	var student dto.StudentProfileResponse
	err := r.db.Get(&student, query, studentID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &student, err
}

func (r *studentRepository) GetStudentsByGroup(groupID int) ([]dto.GroupStudentResponse, error) {
	query := `
        SELECT 
            s.id,
            u.name,
            s.student_id_number,
            u.email
        FROM students s
        JOIN users u ON s.user_id = u.id
        WHERE s.group_id = $1
        ORDER BY u.name
    `

	var students []dto.GroupStudentResponse
	err := r.db.Select(&students, query, groupID)
	return students, err
}
