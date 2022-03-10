package controllers

import (
	"fmt"

	"github.com/achimonchi/pkg/models"
)

type UserController interface {
	SayHello()
	GetPhone()
}

type user struct {
	User *models.User
}

func NewUserController(model *models.User) UserController {
	return &user{
		User: model,
	}
}

func (u *user) SayHello() {
	fmt.Println("Hello, ", u.User.Name)
}

func (u *user) GetPhone() {
	fmt.Println("My phone is :", u.User.Phone)
}
