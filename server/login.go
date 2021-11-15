package server

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"myServer/config"
)

// 不在这儿作比较，用username找出并返回password
//func login (username string, password string) (bool, error) {
//	p, err := getPassword(username)
//	if err != nil{
//		return false, err
//	}
//	if p != password{
//		return false, nil
//	} else {
//		return true, nil
//	}
//}

func getPassword(username string) ( error){
	db := config.GetDB()
	p := Password{}
	tx := db.Where("user_name = ?", username).Take(&p)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound){
		// fmt.Println("No Data Find")
		return tx.Error
	} else if tx.Error != nil{
		// fmt.Println("Query failed", tx.Error)
		return  tx.Error
	}
	log.Printf("id: %v, username: %v,  pssword: %v, create_date: %v \n",p.Id, p.UserName, p.Password, p.CreateAt)
	return nil
}