package sql

import "time"

type Password struct {
	Id       int `gorm:"column:id;PRIMARY_KEY"`
	UserName string
	Password string
	CreateAt time.Time `gorm:"column:create_date"`
	UpdateAt time.Time `gorm:"column:update_date"`
}

type Directory struct {
	FileName   string    `gorm:"column:FileName"`
	UserName   string    `gorm:"column:UserName"`
	FileType   string    `gorm:"column:FileType"`
	UploadTime time.Time `gorm:"column:UploadTime;default:null"`
	Id         int       `gorm:"column:Id; PRIMARY_KEY"`
}
