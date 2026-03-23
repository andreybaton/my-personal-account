package auth

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	// поменять NOW() на DEFAULT NOW() в бд
	query := `
        INSERT INTO users (email, password, first_name, last_name, role, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `

	return r.db.QueryRow(
		query,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Role,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, password, first_name, last_name, role, created_at, updated_at 
              FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password,
		&user.FirstName, &user.LastName, &user.Role,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByID(id int) (*User, error) {
	user := &User{}
	query := `SELECT id, email, first_name, last_name, role, created_at, updated_at 
              FROM users WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email,
		&user.FirstName, &user.LastName, &user.Role,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}
