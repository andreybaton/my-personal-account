package service

import (
	"fmt"
	"time"

	"backend/internal/dto"
	"backend/internal/entity"
	"backend/internal/repository"
)

type attendanceService struct {
	attendanceRepo repository.AttendanceRepository
}

func NewAttendanceService(attendanceRepo repository.AttendanceRepository) AttendanceService {
	return &attendanceService{attendanceRepo: attendanceRepo}
}

func (s *attendanceService) MarkAttendance(request dto.MarkAttendanceRequest) error {
	attendance := &entity.Attendance{
		StudentID: request.StudentID,
		LessonID:  request.LessonID,
		Status:    request.Status,
	}

	return s.attendanceRepo.MarkAttendance(attendance)
}

func (s *attendanceService) MarkAttendanceBatch(requests []dto.MarkAttendanceRequest) error {
	attendances := make([]entity.Attendance, len(requests))
	for i, req := range requests {
		attendances[i] = entity.Attendance{
			StudentID: req.StudentID,
			LessonID:  req.LessonID,
			Status:    req.Status,
		}
	}

	return s.attendanceRepo.MarkAttendanceBatch(attendances)
}

func (s *attendanceService) GetStudentAttendance(studentID int, startDate, endDate string) ([]dto.AttendanceResponse, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	return s.attendanceRepo.GetStudentAttendanceWithDetails(studentID, start, end)
}

func (s *attendanceService) GetAttendanceStats(studentID int, period string) (*dto.AttendanceStatsResponse, error) {
	return s.attendanceRepo.GetAttendanceStatsByPeriod(studentID, period)
}

func (s *attendanceService) GetLessonAttendance(lessonID int) (interface{}, error) {
	return s.attendanceRepo.GetLessonAttendanceWithDetails(lessonID)
}

func (s *attendanceService) GetGroupAttendanceSummary(groupID int, startDate, endDate string) (interface{}, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	return s.attendanceRepo.GetAttendanceSummaryByGroup(groupID, start, end)
}
