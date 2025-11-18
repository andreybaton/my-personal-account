package handler

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ScheduleHandler struct {
	scheduleService service.ScheduleService
	validator       *validator.Validate
}

func NewScheduleHandler(scheduleService service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: scheduleService,
		validator:       validator.New(),
	}
}

// GetStudentSchedule godoc
// @Summary Получить расписание студента
// @Description Возвращает расписание студента на указанный период
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "ID студента" example(1)
// @Param start_date query string false "Начальная дата (формат: 2006-01-02)" example("2024-01-01")
// @Param end_date query string false "Конечная дата (формат: 2006-01-02)" example("2024-01-31")
// @Success 200 {array} dto.LessonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/students/{id} [get]
func (h *ScheduleHandler) GetStudentSchedule(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	schedule, err := h.scheduleService.GetStudentSchedule(studentID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// GetWeeklySchedule godoc
// @Summary Получить недельное расписание студента
// @Description Возвращает расписание студента на указанную неделю
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "ID студента" example(1)
// @Param week_start query string false "Начало недели (формат: 2006-01-02)" example("2024-01-01")
// @Success 200 {object} dto.WeeklyScheduleResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/students/{id}/weekly [get]
func (h *ScheduleHandler) GetWeeklySchedule(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	weekStart := c.DefaultQuery("week_start", "")

	schedule, err := h.scheduleService.GetWeeklySchedule(studentID, weekStart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// GetDailySchedule godoc
// @Summary Получить дневное расписание студента
// @Description Возвращает расписание студента на указанный день
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "ID студента" example(1)
// @Param date query string false "Дата (формат: 2006-01-02)" example("2024-01-01")
// @Success 200 {object} dto.DayScheduleResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/students/{id}/daily [get]
func (h *ScheduleHandler) GetDailySchedule(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	date := c.DefaultQuery("date", "")

	schedule, err := h.scheduleService.GetDailySchedule(studentID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// GetGroupSchedule godoc
// @Summary Получить расписание группы
// @Description Возвращает расписание группы на указанный период
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "ID группы" example(1)
// @Param start_date query string false "Начальная дата (формат: 2006-01-02)" example("2024-01-01")
// @Param end_date query string false "Конечная дата (формат: 2006-01-02)" example("2024-01-31")
// @Success 200 {array} dto.LessonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/groups/{id} [get]
func (h *ScheduleHandler) GetGroupSchedule(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid group ID"})
		return
	}

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	schedule, err := h.scheduleService.GetGroupSchedule(groupID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// GetTeacherSchedule godoc
// @Summary Получить расписание преподавателя
// @Description Возвращает расписание преподавателя на указанный период
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "ID преподавателя" example(1)
// @Param start_date query string false "Начальная дата (формат: 2006-01-02)" example("2024-01-01")
// @Param end_date query string false "Конечная дата (формат: 2006-01-02)" example("2024-01-31")
// @Success 200 {array} dto.LessonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/teachers/{id} [get]
func (h *ScheduleHandler) GetTeacherSchedule(c *gin.Context) {
	teacherID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid teacher ID"})
		return
	}

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	schedule, err := h.scheduleService.GetTeacherSchedule(teacherID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// GetClassroomSchedule godoc
// @Summary Получить расписание аудитории
// @Description Возвращает расписание аудитории на указанный период
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "ID аудитории" example(1)
// @Param start_date query string false "Начальная дата (формат: 2006-01-02)" example("2024-01-01")
// @Param end_date query string false "Конечная дата (формат: 2006-01-02)" example("2024-01-31")
// @Success 200 {array} dto.LessonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/classrooms/{id} [get]
func (h *ScheduleHandler) GetClassroomSchedule(c *gin.Context) {
	classroomID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid classroom ID"})
		return
	}

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	schedule, err := h.scheduleService.GetClassroomSchedule(classroomID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// CreateLesson godoc
// @Summary Создать занятие
// @Description Создает новое занятие в расписании
// @Tags schedule
// @Accept json
// @Produce json
// @Param request body dto.CreateLessonRequest true "Данные для создания занятия"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/lessons [post]
func (h *ScheduleHandler) CreateLesson(c *gin.Context) {
	var request dto.CreateLessonRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if err := h.validator.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.scheduleService.CreateLesson(request); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse{Message: "Lesson created successfully"})
}

// UpdateLesson godoc
// @Summary Обновить занятие
// @Description Обновляет существующее занятие в расписании
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "ID занятия" example(1)
// @Param request body dto.CreateLessonRequest true "Данные для обновления занятия"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/lessons/{id} [put]
func (h *ScheduleHandler) UpdateLesson(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid lesson ID"})
		return
	}

	var request dto.CreateLessonRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if err := h.validator.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.scheduleService.UpdateLesson(lessonID, request); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{Message: "Lesson updated successfully"})
}

// DeleteLesson godoc
// @Summary Удалить занятие
// @Description Удаляет занятие из расписания
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "ID занятия" example(1)
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/lessons/{id} [delete]
func (h *ScheduleHandler) DeleteLesson(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid lesson ID"})
		return
	}

	if err := h.scheduleService.DeleteLesson(lessonID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{Message: "Lesson deleted successfully"})
}

// GetLessonDetails godoc
// @Summary Получить детали занятия
// @Description Возвращает детальную информацию о занятии
// @Tags schedule
// @Accept json
// @Produce json
// @Param id path int true "ID занятия" example(1)
// @Success 200 {object} dto.LessonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /schedule/lessons/{id} [get]
func (h *ScheduleHandler) GetLessonDetails(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid lesson ID"})
		return
	}

	lesson, err := h.scheduleService.GetLessonDetails(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if lesson == nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Lesson not found"})
		return
	}

	c.JSON(http.StatusOK, lesson)
}
