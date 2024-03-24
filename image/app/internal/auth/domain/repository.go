package domain

type UserRepository interface {
	Save(user *User)
	GetByEmail(email string) (*User, error)
	GetPassHashByEmail(email string) (string, error)
}
