package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myServer/utils"
)

var DB *gorm.DB

func initDb(){
	dsn := getDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		utils.ErrorRecorder(err)
		return
	}
	DB = db
}

func getDsn () string{
	userName := Config.GetString("userName")
	password := Config.GetString("password")
	protocol := Config.GetString("protocol")
	address := Config.GetString("address")
	dbName := Config.GetString("dbName")
	paramValue := Config.GetString("paramValue")

	return userName+":"+password+"@"+protocol+"("+address+")/"+dbName+"?"+paramValue+"charset=utf8&parseTime=True"
}

func GetDB() *gorm.DB {
	db := DB.Session(&gorm.Session{NewDB: true})
	return db
}
