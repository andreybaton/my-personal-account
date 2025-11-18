package handler

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
	}
}

// GetStudentProfile godoc
// @Summary Получить профиль студента
// @Description Возвращает профиль студента с основной информацией
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "ID студента" example(1)
// @Success 200 {object} dto.StudentProfileResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /students/{id}/profile [get]
func (h *StudentHandler) GetStudentProfile(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	profile, err := h.studentService.GetStudentProfile(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if profile == nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Student not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// GetStudentAcademicInfo godoc
// @Summary Получить академическую информацию
// @Description Возвращает академическую информацию о студенте (оценки, посещаемость и т.д.)
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "ID студента" example(1)
// @Success 200 {object} dto.StudentAcademicResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /students/{id}/academic [get]
func (h *StudentHandler) GetStudentAcademicInfo(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid student ID"})
		return
	}

	academicInfo, err := h.studentService.GetStudentAcademicInfo(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, academicInfo)
}
