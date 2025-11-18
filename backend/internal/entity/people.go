package entity

type Employee struct {
	ID       int    `db:"id"`
	UserID   int    `db:"user_id"`
	Position string `db:"position"`
}

type Teacher struct {
	ID             int    `db:"id"`
	AcademicDegree string `db:"academic_degree"`
	DepartmentID   int    `db:"department_id"`
	EmployeeID     int    `db:"employee_id"`
}

type Student struct {
	ID              int    `db:"id"`
	StudentIDNumber string `db:"student_id_number"`
	GroupID         *int   `db:"group_id"` // может быть NULL
	UserID          int    `db:"user_id"`
}
