package application

import (
	"api/internal/auth/application/commands"
	"context"
)

type AuthModule struct {
	Commands *commands.Commands
}

func NewModule(ctx context.Context) *AuthModule {
	return &AuthModule{
		Commands: commands.NewCommands(ctx),
	}
}
