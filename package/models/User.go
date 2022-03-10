package models

type User struct {
	Name  string
	Phone string
}

func NewUser(name string, phone string) *User {
	return &User{
		Name:  name,
		Phone: phone,
	}
}
