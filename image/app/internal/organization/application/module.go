package application

import (
	"api/internal/organization/application/commands"
	"api/internal/organization/application/queries"
	"context"
)

type OrganizationModule struct {
	Commands *commands.Commands
	Queries  *queries.Queries
}

func NewModule(ctx context.Context) *OrganizationModule {
	return &OrganizationModule{
		commands.NewCommands(ctx),
		queries.NewQueries(ctx),
	}
}
