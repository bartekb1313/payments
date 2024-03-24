package application

import (
	"api/internal/auth/domain"
	"api/internal/auth/infrastructure/persistence/sql"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type UserCommands struct {
	ctx        context.Context
	repository domain.UserRepository
}

func (s *UserCommands) CreateUser(name string, email string, password string) domain.User {
	user := domain.NewUser(name, email, s.HashPassword(password))
	s.repository.Save(&user)
	return user
}

func (s *UserCommands) HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func (s *UserCommands) CheckPassword(email, password string) bool {
	hashedPass, _ := s.repository.GetPassHashByEmail(email)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return err == nil
}

func NewUserServices(ctx context.Context) *UserCommands {
	return &UserCommands{
		ctx:        ctx,
		repository: sql.NewUserRepository(&ctx),
	}
}
