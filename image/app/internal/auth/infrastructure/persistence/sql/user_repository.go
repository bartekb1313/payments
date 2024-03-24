package sql

import (
	"api/internal/auth/domain"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type Repo struct {
	ctx    *context.Context
	dbpool *pgxpool.Pool
}

func (r Repo) GetPassHashByEmail(email string) (string, error) {
	var hashedPass string
	r.dbpool.QueryRow(context.Background(), "select password_hash from users where email = ($1)", email).Scan(&hashedPass)
	return hashedPass, nil
}

func (r Repo) Save(user *domain.User) {
	_, err := r.dbpool.Exec(context.Background(), "insert into users(name, email, password_hash) values($1, $2, $3)", user.GetName(), user.GetEmail(), user.GetPassword())
	if err != nil {
		fmt.Println(err)
	}
}

func (r Repo) GetByEmail(email string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(ctx *context.Context) *Repo {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return &Repo{
		ctx,
		dbpool,
	}
}
