package service

import (
	"backend/internal/dto"
	"backend/internal/repository"
	"fmt"
	"time"
)

type studentService struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &studentService{studentRepo: studentRepo}
}

func (s *studentService) GetStudentProfile(studentID int) (*dto.StudentProfileResponse, error) {
	return s.studentRepo.GetStudentWithDetails(studentID)
}

func (s *studentService) GetStudentAcademicInfo(studentID int) (*dto.StudentAcademicResponse, error) {
	profile, err := s.studentRepo.GetStudentWithDetails(studentID)
	if err != nil {
		return nil, err
	}

	currentSemester := "1"
	if profile.AdmissionYear > 0 {
		currentYear := time.Now().Year()
		yearsPassed := currentYear - profile.AdmissionYear
		currentSemester = fmt.Sprintf("%d", yearsPassed*2+1) // 1 семестр в год
	}

	return &dto.StudentAcademicResponse{
		StudentID:        studentID,
		AverageGrade:     nil,                   // TODO: реализовать расчет среднего балла
		TotalDisciplines: 0,                     // TODO: реализовать подсчет дисциплин
		CompletedCredits: 0,                     // TODO: реализовать подсчет кредитов
		CurrentSemester:  currentSemester,       // Используем вычисленный семестр
		Grades:           []dto.GradeResponse{}, // TODO: реализовать получение оценок
	}, nil
}
