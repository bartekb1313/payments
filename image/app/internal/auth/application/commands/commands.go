package commands

import (
	"api/internal/auth/adapters"
	"api/internal/auth/domain"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type Commands struct {
	ctx        context.Context
	repository domain.UserRepository
}

func (s *Commands) CreateUser(email string, password string) domain.User {
	user := domain.NewUser(email, s.HashPassword(password))
	s.repository.Save(&user)
	return user
}

func (s *Commands) HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func (s *Commands) CheckPassword(email, password string) bool {
	hashedPass, _ := s.repository.GetPassHashByEmail(email)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return err == nil
}

func NewCommands(ctx context.Context) *Commands {
	return &Commands{
		ctx:        ctx,
		repository: adapters.NewUserRepository(&ctx),
	}
}
