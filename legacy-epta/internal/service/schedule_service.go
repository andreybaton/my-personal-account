package service

import (
	"fmt"
	"time"

	"backend/internal/dto"
	"backend/internal/entity"
	"backend/internal/repository"
)

type scheduleService struct {
	scheduleRepo repository.ScheduleRepository
}

func NewScheduleService(scheduleRepo repository.ScheduleRepository) ScheduleService {
	return &scheduleService{scheduleRepo: scheduleRepo}
}

func (s *scheduleService) GetWeeklySchedule(studentID int, weekStart string) (*dto.WeeklyScheduleResponse, error) {
	startDate, err := time.Parse("2006-01-02", weekStart)
	if err != nil {
		return nil, fmt.Errorf("invalid week start date: %v", err)
	}

	_, week := startDate.ISOWeek()
	startOfWeek := getStartOfWeek(startDate)
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	lessons, err := s.scheduleRepo.GetStudentScheduleWithDetails(studentID, startOfWeek, endOfWeek)
	if err != nil {
		return nil, err
	}

	days := s.groupLessonsByDay(lessons, startOfWeek)

	return &dto.WeeklyScheduleResponse{
		WeekNumber: week,
		StartDate:  startOfWeek.Format("2006-01-02"),
		EndDate:    endOfWeek.Format("2006-01-02"),
		Days:       days,
	}, nil
}

func (s *scheduleService) GetDailySchedule(studentID int, date string) (*dto.DayScheduleResponse, error) {
	scheduleDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %v", err)
	}

	lessons, err := s.scheduleRepo.GetStudentScheduleWithDetails(studentID, scheduleDate, scheduleDate)
	if err != nil {
		return nil, err
	}

	dayName := getRussianDayName(scheduleDate.Weekday())

	return &dto.DayScheduleResponse{
		Date:    date,
		DayName: dayName,
		Lessons: lessons,
	}, nil
}

func (s *scheduleService) GetStudentSchedule(studentID int, startDate, endDate string) ([]dto.LessonResponse, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	return s.scheduleRepo.GetStudentScheduleWithDetails(studentID, start, end)
}

func (s *scheduleService) GetGroupSchedule(groupID int, startDate, endDate string) ([]dto.LessonResponse, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	return s.scheduleRepo.GetGroupSchedule(groupID, start, end)
}

func (s *scheduleService) GetTeacherSchedule(teacherID int, startDate, endDate string) ([]dto.LessonResponse, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	return s.scheduleRepo.GetTeacherSchedule(teacherID, start, end)
}

func (s *scheduleService) GetClassroomSchedule(classroomID int, startDate, endDate string) ([]dto.LessonResponse, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	return s.scheduleRepo.GetClassroomSchedule(classroomID, start, end)
}

func (s *scheduleService) CreateLesson(request dto.CreateLessonRequest) error {
	if err := s.ValidateLessonConflict(request); err != nil {
		return err
	}

	lessonDate, err := time.Parse("2006-01-02", request.LessonDate)
	if err != nil {
		return fmt.Errorf("invalid lesson date: %v", err)
	}

	for _, groupID := range request.GroupIDs {
		lesson := &entity.Lesson{
			TeacherID:    request.TeacherID,
			ClassroomID:  request.ClassroomID,
			DisciplineID: request.DisciplineID,
			GroupID:      groupID,
			LessonDate:   lessonDate,
			StartTime:    request.StartTime,
			EndTime:      request.EndTime,
			Type:         request.Type,
		}

		if err := s.scheduleRepo.CreateLesson(lesson); err != nil {
			return err
		}
	}

	return nil
}

func (s *scheduleService) UpdateLesson(lessonID int, request dto.CreateLessonRequest) error {
	currentLesson, err := s.scheduleRepo.GetLessonByID(lessonID)
	if err != nil {
		return err
	}
	if currentLesson == nil {
		return fmt.Errorf("lesson not found")
	}

	conflictRequest := request
	excludeID := &lessonID
	if err := s.validateLessonConflictWithExclusion(conflictRequest, excludeID); err != nil {
		return err
	}

	lessonDate, err := time.Parse("2006-01-02", request.LessonDate)
	if err != nil {
		return fmt.Errorf("invalid lesson date: %v", err)
	}

	// предполагаем, что занятие для одной группы
	if len(request.GroupIDs) > 0 {
		currentLesson.GroupID = request.GroupIDs[0]
	}
	currentLesson.TeacherID = request.TeacherID
	currentLesson.ClassroomID = request.ClassroomID
	currentLesson.DisciplineID = request.DisciplineID
	currentLesson.LessonDate = lessonDate
	currentLesson.StartTime = request.StartTime
	currentLesson.EndTime = request.EndTime
	currentLesson.Type = request.Type

	return s.scheduleRepo.UpdateLesson(currentLesson)
}

func (s *scheduleService) DeleteLesson(lessonID int) error {
	return s.scheduleRepo.DeleteLesson(lessonID)
}

func (s *scheduleService) ValidateLessonConflict(request dto.CreateLessonRequest) error {
	return s.validateLessonConflictWithExclusion(request, nil)
}

func (s *scheduleService) validateLessonConflictWithExclusion(request dto.CreateLessonRequest, excludeLessonID *int) error {
	lessonDate, err := time.Parse("2006-01-02", request.LessonDate)
	if err != nil {
		return fmt.Errorf("invalid lesson date: %v", err)
	}

	// проверка конфликта для преподавателя
	hasConflict, err := s.scheduleRepo.CheckTeacherConflict(
		request.TeacherID,
		lessonDate,
		request.StartTime,
		request.EndTime,
		excludeLessonID,
	)
	if err != nil {
		return err
	}
	if hasConflict {
		return fmt.Errorf("teacher has conflicting schedule")
	}

	// проверка конфликта для аудитории
	hasConflict, err = s.scheduleRepo.CheckClassroomConflict(
		request.ClassroomID,
		lessonDate,
		request.StartTime,
		request.EndTime,
		excludeLessonID,
	)
	if err != nil {
		return err
	}
	if hasConflict {
		return fmt.Errorf("classroom is occupied")
	}

	// проверка конфликта для каждой группы
	for _, groupID := range request.GroupIDs {
		hasConflict, err = s.scheduleRepo.CheckGroupConflict(
			groupID,
			lessonDate,
			request.StartTime,
			request.EndTime,
			excludeLessonID,
		)
		if err != nil {
			return err
		}
		if hasConflict {
			return fmt.Errorf("group %d has conflicting schedule", groupID)
		}
	}

	return nil
}

func (s *scheduleService) GetLessonDetails(lessonID int) (*dto.LessonResponse, error) {
	return s.scheduleRepo.GetLessonWithDetails(lessonID)
}

func (s *scheduleService) groupLessonsByDay(lessons []dto.LessonResponse, startOfWeek time.Time) []dto.DayScheduleResponse {
	days := make([]dto.DayScheduleResponse, 7)

	for i := 0; i < 7; i++ {
		currentDate := startOfWeek.AddDate(0, 0, i)
		dayName := getRussianDayName(currentDate.Weekday())

		dayLessons := []dto.LessonResponse{}
		for _, lesson := range lessons {
			if lesson.Date == currentDate.Format("2006-01-02") {
				dayLessons = append(dayLessons, lesson)
			}
		}

		days[i] = dto.DayScheduleResponse{
			Date:    currentDate.Format("2006-01-02"),
			DayName: dayName,
			Lessons: dayLessons,
		}
	}

	return days
}
