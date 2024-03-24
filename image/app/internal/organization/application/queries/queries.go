package queries

import (
	"api/internal/organization/adapters"
	"api/internal/organization/domain"
	"context"
)

type BranchQueries struct {
	ctx        context.Context
	repository domain.BranchRepository
}

func (s *BranchQueries) GetBranches() ([]domain.Branch, error) {
	return s.repository.GetAll()
}

func NewBranchQueries(ctx context.Context) *BranchQueries {
	return &BranchQueries{
		ctx:        ctx,
		repository: adapters.NewBranchRepository(&ctx),
	}
}
