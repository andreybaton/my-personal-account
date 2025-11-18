package handler

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	groupService service.GroupService
}

func NewGroupHandler(groupService service.GroupService) *GroupHandler {
	return &GroupHandler{
		groupService: groupService,
	}
}

// GetAllGroups godoc
// @Summary Получить все группы
// @Description Возвращает список всех групп
// @Tags groups
// @Accept json
// @Produce json
// @Success 200 {array} dto.GroupResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups [get]
func (h *GroupHandler) GetAllGroups(c *gin.Context) {
	groups, err := h.groupService.GetAllGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GetGroupDetails godoc
// @Summary Получить детали группы
// @Description Возвращает детальную информацию о группе включая студентов
// @Tags groups
// @Accept json
// @Produce json
// @Param id path int true "ID группы" example(1)
// @Success 200 {object} dto.GroupDetailsResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups/{id} [get]
func (h *GroupHandler) GetGroupDetails(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid group ID"})
		return
	}

	group, err := h.groupService.GetGroupDetails(groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if group == nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Group not found"})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetGroupsByFaculty godoc
// @Summary Получить группы по факультету
// @Description Возвращает список групп по ID факультета
// @Tags groups
// @Accept json
// @Produce json
// @Param faculty_id path int true "ID факультета" example(1)
// @Success 200 {array} dto.GroupResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups/faculty/{faculty_id} [get]
func (h *GroupHandler) GetGroupsByFaculty(c *gin.Context) {
	facultyID, err := strconv.Atoi(c.Param("faculty_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid faculty ID"})
		return
	}

	groups, err := h.groupService.GetGroupsByFaculty(facultyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}
