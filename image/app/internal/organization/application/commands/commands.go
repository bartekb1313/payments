package commands

import (
	"api/internal/organization/adapters"
	"api/internal/organization/domain"
	"context"
)

type Commands struct {
	ctx        context.Context
	repository domain.BranchRepository
}

func NewCommands(ctx context.Context) *Commands {
	return &Commands{
		ctx:        ctx,
		repository: adapters.NewBranchRepository(&ctx),
	}

}
