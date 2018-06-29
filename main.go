package main

import (
	"fmt"
	"pearbit/MySQL_Crud_Desktop/controller"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Application Starts . . .")

	//user := models.UserStruct{Username: "unique", Password: "123", Enable: 1, UserID: 0}

	//controller.Create(user)

	user := controller.Login("unique", "123")

	fmt.Println("Success login for: ", user.Username)
}
