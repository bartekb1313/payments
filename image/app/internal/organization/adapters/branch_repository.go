package adapters

import (
	"api/internal/organization/domain"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type Repo struct {
	ctx    *context.Context
	dbpool *pgxpool.Pool
}

func (repo Repo) Save(branch *domain.Branch) {
	fmt.Println("SAVE")
	_, err := repo.dbpool.Exec(context.Background(), "insert into branches(name) values($1)", branch.Name())
	if err != nil {
		fmt.Println(err)
	}
}

func (repo Repo) GetAll() ([]domain.Branch, error) {
	rows, err := repo.dbpool.Query(context.Background(), "select name from branches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branches []domain.Branch
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		branch := domain.NewBranch(name)
		branches = append(branches, branch)
	}
	return branches, nil
}

func NewBranchRepository(ctx *context.Context) *Repo {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &Repo{
		ctx,
		dbpool,
	}
}
