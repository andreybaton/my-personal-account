package service

import (
	"backend/internal/dto"
	"backend/internal/repository"
)

type groupService struct {
	groupRepo repository.GroupRepository
}

func NewGroupService(groupRepo repository.GroupRepository) GroupService {
	return &groupService{groupRepo: groupRepo}
}

func (s *groupService) GetGroupDetails(groupID int) (*dto.GroupDetailsResponse, error) {
	return s.groupRepo.GetGroupWithDetails(groupID)
}

func (s *groupService) GetGroupsByFaculty(facultyID int) ([]dto.GroupResponse, error) {
	return s.groupRepo.GetGroupsByFaculty(facultyID)
}

func (s *groupService) GetAllGroups() ([]dto.GroupResponse, error) {
	return s.groupRepo.GetAllGroups()
}
