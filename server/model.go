package server

import "time"

type Password struct {
	Id       int `gorm:"column:id; PRIMARY_KEY"`
	UserName string
	Password string
	CreateAt time.Time `gorm:"column:create_date"`
	UpdateAt time.Time `gorm:"column:update_date"`
}

