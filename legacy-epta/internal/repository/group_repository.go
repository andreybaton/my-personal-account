package repository

import (
	"database/sql"

	"backend/internal/dto"
	"backend/internal/entity"

	"github.com/jmoiron/sqlx"
)

type groupRepository struct {
	db *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) GetGroupByID(groupID int) (*entity.Group, error) {
	query := `SELECT * FROM groups WHERE id = $1`

	var group entity.Group
	err := r.db.Get(&group, query, groupID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &group, err
}

func (r *groupRepository) GetGroupWithDetails(groupID int) (*dto.GroupDetailsResponse, error) {
	groupQuery := `
        SELECT 
            g.id,
            g.name,
            g.admission_year,
            g.specialization,
            f.name as faculty,
            COUNT(s.id) as student_count
        FROM groups g
        JOIN faculties f ON g.faculty_id = f.id
        LEFT JOIN students s ON s.group_id = g.id
        WHERE g.id = $1
        GROUP BY g.id, g.name, g.admission_year, g.specialization, f.name
    `

	var group dto.GroupResponse
	err := r.db.Get(&group, groupQuery, groupID)
	if err != nil {
		return nil, err
	}

	students, err := r.GetGroupStudents(groupID)
	if err != nil {
		return nil, err
	}

	return &dto.GroupDetailsResponse{
		GroupResponse: group,
		Students:      students,
	}, nil
}

func (r *groupRepository) GetGroupsByFaculty(facultyID int) ([]dto.GroupResponse, error) {
	query := `
        SELECT 
            g.id,
            g.name,
            g.admission_year,
            g.specialization,
            f.name as faculty,
            COUNT(s.id) as student_count
        FROM groups g
        JOIN faculties f ON g.faculty_id = f.id
        LEFT JOIN students s ON s.group_id = g.id
        WHERE g.faculty_id = $1
        GROUP BY g.id, g.name, g.admission_year, g.specialization, f.name
        ORDER BY g.admission_year DESC, g.name
    `

	var groups []dto.GroupResponse
	err := r.db.Select(&groups, query, facultyID)
	return groups, err
}

func (r *groupRepository) GetAllGroups() ([]dto.GroupResponse, error) {
	query := `
        SELECT 
            g.id,
            g.name,
            g.admission_year,
            g.specialization,
            f.name as faculty,
            COUNT(s.id) as student_count
        FROM groups g
        JOIN faculties f ON g.faculty_id = f.id
        LEFT JOIN students s ON s.group_id = g.id
        GROUP BY g.id, g.name, g.admission_year, g.specialization, f.name
        ORDER BY g.admission_year DESC, g.name
    `

	var groups []dto.GroupResponse
	err := r.db.Select(&groups, query)
	return groups, err
}

func (r *groupRepository) GetGroupStudents(groupID int) ([]dto.GroupStudentResponse, error) {
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
