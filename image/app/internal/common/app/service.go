package app

import (
	"api/internal/auth/application"
	"api/internal/organization/application/commands"
	"api/internal/organization/application/queries"
	"context"
)

type Application struct {
	BranchCommands *commands.BranchCommands
	BranchQueries  *queries.BranchQueries
	UserCommands   *application.UserCommands
}

func NewApplication(ctx context.Context) *Application {
	return &Application{
		BranchCommands: commands.NewBranchServices(ctx),
		BranchQueries:  queries.NewBranchQueries(ctx),
		UserCommands:   application.NewUserServices(ctx),
	}
}
