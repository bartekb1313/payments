package commands

import (
	"api/internal/organization/domain"
)

func (s *BranchCommands) CreateBranch(name string) domain.Branch {
	branch := domain.NewBranch(name)
	s.repository.Save(&branch)
	return domain.NewBranch(name)
}
