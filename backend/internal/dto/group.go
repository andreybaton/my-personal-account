package dto

type GroupResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	AdmissionYear  int    `json:"admission_year"`
	Specialization string `json:"specialization"`
	Faculty        string `json:"faculty"`
	StudentCount   int    `json:"student_count"`
}

type GroupDetailsResponse struct {
	GroupResponse
	Students []GroupStudentResponse `json:"students"`
}

type GroupStudentResponse struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	StudentIDNumber string `json:"student_id_number"`
	Email           string `json:"email"`
}
