package domain

type BranchRepository interface {
	Save(branch *Branch)
	GetAll() ([]Branch, error)
}
