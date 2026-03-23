package dto

type StudentProfileResponse struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	StudentIDNumber string `json:"student_id_number"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	GroupID         *int   `json:"group_id,omitempty"`
	GroupName       string `json:"group_name,omitempty"`
	FacultyID       int    `json:"faculty_id"`
	FacultyName     string `json:"faculty_name"`
	Specialization  string `json:"specialization"`
	AdmissionYear   int    `json:"admission_year"`
}

type StudentAcademicResponse struct {
	StudentID        int             `json:"student_id"`
	AverageGrade     *float64        `json:"average_grade,omitempty"`
	TotalDisciplines int             `json:"total_disciplines"`
	CompletedCredits int             `json:"completed_credits"`
	CurrentSemester  string          `json:"current_semester"`
	Grades           []GradeResponse `json:"grades"`
}

type GradeResponse struct {
	ID           int      `json:"id"`
	DisciplineID int      `json:"discipline_id"`
	Discipline   string   `json:"discipline"`
	GradeValue   *float64 `json:"grade_value,omitempty"`
	GradeType    string   `json:"grade_type"`
	Semester     string   `json:"semester"`
	CreatedAt    string   `json:"created_at"`
}
