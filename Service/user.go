package Service

import (
	"log"
	"myServer/config"
	"time"
)

type UserInfo struct {
	userID      uint      `gorm:"userID"`
	userName    string    `gorm:"size:25;not null"`
	password    string    `gorm:"size:25;not null"`
	createdDate time.Time `gorm:"created_date"`
}

func SetUserInfo(userName, password string, date time.Time) *UserInfo {

	return &UserInfo{userName: userName, password: password, createdDate: date}
}

func (u UserInfo) Login() bool {
	db := config.GetDB()

	var p UserInfo
	db.Table("userinfo").Where("userID = ?", "1").Find(&p)
	// log.Printf("p%v\n", p)
	//
	// db.Table("userinfo").Where("userName = ?", u.UserName).Take(&p)
	// // 对比一下密码？
	log.Printf("ID: %v, name: %v, password: %v, Date: %v\n", p.userID, p.userName, p.password, p.createdDate)
	return false
}

func (u UserInfo) Insert() bool {
	db := config.GetDB()
	// var insert UserInfo
	err := db.Table("userinfo").Create(&UserInfo{}).Error
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	return false
}
