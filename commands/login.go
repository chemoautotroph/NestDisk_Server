package commands

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"myServer/config"
	"myServer/sql"
)

var loggedIn bool

func Login(username, password string) (string, error, string) {
	p, err, errorMessage:= getPassword(username)
	if err != nil {
		return "invalid username", err, errorMessage
	}
	if p == password{
		loggedIn = true
		return "Login Successful", nil, errorMessage
	} else {
		return "invalid password",nil, errorMessage
	}
}

func getPassword(username string) (string, error, string) {
	db := config.GetDB()
	p := sql.Password{}
	tx := db.Where("user_name = ?", username).Take(&p)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return "", tx.Error, "No Data Find"
	} else if tx.Error != nil {
		return "", tx.Error, "Query failed"
	}
	log.Printf("id: %v, username: %v,  pssword: %v, create_date: %v \n", p.Id, p.UserName, p.Password, p.CreateAt)
	return p.Password, nil, ""
}

func GetLoggedIn() bool{
	return loggedIn
}