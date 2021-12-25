package commands

import (
	"fmt"
	"myServer/config"
	"myServer/sql"
	"time"
)

func Register(username, password string) error {
	p := sql.Password{
		UserName: username,
		Password: password,
		CreateAt: time.Now(),
	}
	db := config.GetDB()
	db.Create(&p)
	if err := db.Create(&p).Error; err != nil{
		fmt.Println("fail to register")
		return err
	}
	return nil
}
