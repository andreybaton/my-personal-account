package handler

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AttendanceHandler struct {
	attendanceService service.AttendanceService
	validator         *validator.Validate
}

func NewAttendanceHandler(attendanceService service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
		validator:         validator.New(),
	}
}

// MarkAttendance godoc
// @Summary Отметить посещение
// @Description Отмечает посещение студента на занятии
// @Tags attendance
// @Accept json
// @Produce json
// @Param request body dto.MarkAttendanceRequest true "Данные посещения"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /attendance/mark [post]
func (h *AttendanceHandler) MarkAttendance(c *gin.Context) {
	var request dto.MarkAttendanceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if err := h.validator.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.attendanceService.MarkAttendance(request); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{Message: "Attendance marked successfully"})
}

// MarkAttendanceBatch godoc
// @Summary Отметить посещение для нескольких студентов
// @Description Отмечает посещение для нескольких студентов на занятии
// @Tags attendance
// @Accept json
// @Produce json
// @Param request body []dto.MarkAttendanceRequest true "Данные посещений"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /attendance/mark-batch [post]
func (h *AttendanceHandler) MarkAttendanceBatch(c *gin.Context) {
	var requests []dto.MarkAttendanceRequest
	if err := c.ShouldBindJSON(&requests); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body"})
		return
	}

	for _, request := range requests {
		if err := h.validator.Struct(request); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}
	}

	if err := h.attendanceService.MarkAttendanceBatch(requests); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{Message: "Attendance marked successfully for all students"})
}

// GetStudentAttendance godoc
// @Summary Получить посещаемость студента
// @Description Возвращает посещаемость студента за указанный период
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "ID студента" example(1)
// @Param start_date query string false "Начальная дата (формат: 2006-01-02)" example("2024-01-01")
// @Param end_date query string false "Конечная дата (формат: 2006-01-02)" example("2024-01-31")
// @Success 200 {array} dto.AttendanceResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /attendance/students/{id} [get]
func (h *AttendanceHandler) GetStudentAttendance(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	attendance, err := h.attendanceService.GetStudentAttendance(studentID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

// GetAttendanceStats godoc
// @Summary Получить статистику посещаемости
// @Description Возвращает статистику посещаемости студента за период
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "ID студента" example(1)
// @Param period query string false "Период (week, month, semester)" Enums(week, month, semester) default(week)
// @Success 200 {object} dto.AttendanceStatsResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /attendance/students/{id}/stats [get]
func (h *AttendanceHandler) GetAttendanceStats(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	period := c.DefaultQuery("period", "week")

	stats, err := h.attendanceService.GetAttendanceStats(studentID, period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

// GetLessonAttendance godoc
// @Summary Получить посещаемость занятия
// @Description Возвращает посещаемость конкретного занятия
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "ID занятия" example(1)
// @Success 200 {object} dto.LessonAttendanceResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /attendance/lessons/{id} [get]
func (h *AttendanceHandler) GetLessonAttendance(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid lesson ID"})
		return
	}

	attendance, err := h.attendanceService.GetLessonAttendance(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

// GetGroupAttendanceSummary godoc
// @Summary Получить сводку посещаемости по группе
// @Description Возвращает сводку посещаемости всех студентов группы за период
// @Tags attendance
// @Accept json
// @Produce json
// @Param id path int true "ID группы" example(1)
// @Param start_date query string false "Начальная дата (формат: 2006-01-02)" example("2024-01-01")
// @Param end_date query string false "Конечная дата (формат: 2006-01-02)" example("2024-01-31")
// @Success 200 {array} dto.GroupAttendanceSummary
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /attendance/groups/{id}/summary [get]
func (h *AttendanceHandler) GetGroupAttendanceSummary(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid group ID"})
		return
	}

	startDate := c.DefaultQuery("start_date", "")
	endDate := c.DefaultQuery("end_date", "")

	summary, err := h.attendanceService.GetGroupAttendanceSummary(groupID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}
