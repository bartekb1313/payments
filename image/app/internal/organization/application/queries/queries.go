package queries

import (
	"api/internal/organization/adapters"
	"api/internal/organization/domain"
	"context"
)

type Queries struct {
	ctx        context.Context
	repository domain.BranchRepository
}

func (s *Queries) GetBranches() ([]domain.Branch, error) {
	return s.repository.GetAll()
}

func NewQueries(ctx context.Context) *Queries {
	return &Queries{
		ctx:        ctx,
		repository: adapters.NewBranchRepository(&ctx),
	}
}
