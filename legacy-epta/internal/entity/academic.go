package entity

type Building struct {
	ID      int    `db:"id"`
	Type    string `db:"type"`
	Address string `db:"address"`
}

type Faculty struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	BuildingID int    `db:"building_id"`
	DirectorID *int   `db:"director_id"` // может быть NULL
}

type Department struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	FacultyID int    `db:"faculty_id"`
}

type Specialization struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	FacultyID   int    `db:"faculty_id"`
	Description string `db:"description"`
}

type Group struct {
	ID               int    `db:"id"`
	Name             string `db:"name"`
	AdmissionYear    int    `db:"admission_year"`
	SpecializationID int    `db:"specialization_id"`
	FacultyID        int    `db:"faculty_id"`
}
