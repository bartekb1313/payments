package domain

type User struct {
	name     string
	email    string
	password string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func NewUser(name string, email string, password string) User {
	return User{name, email, password}
}
