package app

import (
	auth "api/internal/auth/application"
	organization "api/internal/organization/application"
	"context"
)

type Application struct {
	OrganizationModule *organization.OrganizationModule
	AuthModule         *auth.AuthModule
}

func NewApplication(ctx context.Context) *Application {
	return &Application{
		OrganizationModule: organization.NewModule(ctx),
		AuthModule:         auth.NewModule(ctx),
	}
}
