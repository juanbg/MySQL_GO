package userdao

import (
	"log"
	"pearbit/MySQL_Crud_Desktop/database"
	"pearbit/MySQL_Crud_Desktop/models"

	"golang.org/x/crypto/bcrypt"
)

/*
Create creates a new registry for user table
*/
func Create(user models.UserStruct) bool {
	db := database.DBConn()
	defer db.Close()

	insForm, err := db.Prepare("INSERT INTO users(username, password, enable_status) VALUES(?,?,1)")

	if err != nil {
		panic(err.Error())
	}

	securePass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	insForm.Exec(user.Username, securePass)

	log.Println("INSERT: username: "+user.Username+" | pass: ", string(securePass))

	return true
}

/*
Login perfroms
*/
func Login(user string) models.UserStruct {
	db := database.DBConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM users WHERE username=?", user)

	if err != nil {
		panic(err.Error())
	}

	userStrunct := models.UserStruct{}

	for selDB.Next() {
		var user_id int
		var enable_status int
		var password string
		var username string

		err = selDB.Scan(&user_id, &enable_status, &password, &username)

		if err != nil {
			panic(err.Error())
		}

		userStrunct.Username = username
		userStrunct.Password = password
		userStrunct.UserID = user_id
		userStrunct.Enable = enable_status
	}

	return userStrunct

}
