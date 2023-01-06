package main

import (
	"fmt"
	"go-app/app/models"
)

func main() {

	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@gmail.com"
	u.Password = "testtest"

	u.CreateUser()
}