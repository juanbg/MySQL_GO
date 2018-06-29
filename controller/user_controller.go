package controller

import (
	"log"
	"pearbit/MySQL_Crud_Desktop/dao/user_dao"
	"pearbit/MySQL_Crud_Desktop/models"

	"golang.org/x/crypto/bcrypt"
)

/*
Create controller side
*/
func Create(user models.UserStruct) {
	if user.Username == "" || user.Password == "" {
		log.Fatal("Data Incomplete. not possible")
	}

	userdao.Create(user)
}

/*
Login log in
*/
func Login(username, password string) models.UserStruct {
	if username == "" || password == "" {
		log.Fatal("Data Incomplete. not possible")
	}

	user := userdao.Login(username)

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		log.Fatal("ERROR: User found but password is malformed")
	}

	return user
}
