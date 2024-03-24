package commands

import (
	"api/internal/organization/adapters"
	"api/internal/organization/domain"
	"context"
)

type BranchCommands struct {
	ctx        context.Context
	repository domain.BranchRepository
}

func NewBranchServices(ctx context.Context) *BranchCommands {
	return &BranchCommands{
		ctx:        ctx,
		repository: adapters.NewBranchRepository(&ctx),
	}

}
