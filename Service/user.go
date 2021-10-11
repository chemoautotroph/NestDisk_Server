package Service

import "myServer/config"

type login struct {
	userID   int
	userName string
	password string

}

func Login() {
	db := config.GetDB()

}
