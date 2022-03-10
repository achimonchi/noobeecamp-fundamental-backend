package main

import (
	"github.com/achimonchi/pkg/controllers"
	"github.com/achimonchi/pkg/models"
)

func main() {
	user := models.NewUser("NooBee", "081237583735")

	userController := controllers.NewUserController(user)

	userController.SayHello()
	userController.GetPhone()

}
