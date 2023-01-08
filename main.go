package main

import (
	"fmt"
	"go-app/app/controllers"
	"go-app/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}