package main

import (
	"fmt"
	"go-app/app/models"
)

func main() {

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@gmail.com"
	// u.Password = "testtest"

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()

	// u.Name = "test2"
	// u.Email = "test2@gmail.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@gmail.com"
	// u.Password = "testtest"

	// u.CreateUser()

	u, _ := models.GetUser(2)
	fmt.Println(u)

	u.CreateTodo("掃除")

}