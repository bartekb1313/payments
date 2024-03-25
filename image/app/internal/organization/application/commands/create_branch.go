package commands

import (
	"api/internal/organization/domain"
)

func (s *Commands) CreateBranch(name string) domain.Branch {
	branch := domain.NewBranch(name)
	s.repository.Save(&branch)
	return domain.NewBranch(name)
}
