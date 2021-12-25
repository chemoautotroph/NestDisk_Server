package server

import (
	"fmt"
	"myServer/commands"
	"net"
)

func runCommand (fCommand string, rCommand []string, conn net.Conn) error {
	//fmt.Printf("ping to byte: %v\n", []byte("ping"))
	//fmt.Println("fCommand to byte",[]byte(fCommand))
	login := commands.GetLoggedIn()
	var userName string
	if login == true{
		userName = rCommand[len(rCommand)-1]
		rCommand = rCommand[:len(rCommand)-1]
	}
	switch fCommand {
	case "ping":
		send(conn, []byte("pang"))
		fmt.Println("pang")
	case "login":
		if len(rCommand) != 2 {
			send(conn, []byte("invalid Input."))
			break
		}
		message, err, errorMessage := commands.Login(rCommand[0], rCommand[1])
		if errorMessage != ""{
			send(conn, []byte(errorMessage))
			break
		}
		send(conn, []byte(message))
		fmt.Println("login message: ",message)


		if err != nil {
			return err
		}
	case "register":
		if len(rCommand) != 2 {
			send(conn, []byte("invalid Input."))
			break
		}
		err := commands.Register(rCommand[0], rCommand[1])
		if err != nil {
			fmt.Println(err)
			send(conn, []byte("fail to register"))
			break
		}
		send(conn, []byte("successfully registered"))
	case "upload":
		if userName == ""{
			fmt.Println("please login first")
			send(conn, []byte("please login first, fail to upload"))
			break
		} else {
			fileAddr := rCommand[0]
			fmt.Println("upload with addr ", fileAddr)
			fmt.Println("username is : ", userName)
			err, message:= commands.Upload(fileAddr, userName)
			if err != nil{
				send(conn, []byte("upload error"))
				break
			}
			send(conn, []byte(message))
		}
	case "show":
		if userName == ""{
			fmt.Println("please login first")
			send(conn, []byte("please login first, fail to show"))
			break
		}
		myPrefix := ""
		if len(rCommand) == 1 {
			myPrefix = rCommand[0]
		} else if len(rCommand) > 1 {
			fmt.Println("invalid input")
			send(conn, []byte("invalid input"))
			break
		}
		objectCh, err := commands.Show(userName, myPrefix)
		if err != nil{
			fmt.Println("show err, ", err)
			break
		}
		fmt.Println("All the object stored: ")
		var message string
		for _, s := range objectCh{
			m := fmt.Sprintf("%v\n", s)
			message += m
		}
		send(conn, []byte(message))

	case "download":
		if userName == ""{
			fmt.Println("please login first")
			send(conn, []byte("please login first, fail to download"))
			break
		}
		if len(rCommand) == 2{
			objectName := rCommand[0]
			filePath := rCommand[1]
			err := commands.Download(userName, objectName, filePath)
			if err != nil{
				fmt.Println("commands.Download error ",err)
				send(conn, []byte("Download error"))
				break
			} else {
				message := fmt.Sprintf("Successfully download %v in %v.\n", objectName, filePath)
				send(conn, []byte(message))
			}

		} else {
			send(conn, []byte("invalid input"))
			break
		}


	case "del":

	case "cd":

	case "help":
		if len(rCommand) == 0{
			message := "ping, login, register, upload, download, del, cd. \nFor more information, type help + command."
			send(conn, []byte(message))
			break
		}
		if len(rCommand) != 1{
			message := "please input one command at a time."
			send(conn, []byte(message))
		}

		switch rCommand[0] {
		case "login":
			message := "login username password."
			send(conn, []byte(message))
		case "register":
			message := "register username password."
			send(conn, []byte(message))
		case "upload":
			message := "upload fileAddress"
			send(conn, []byte(message))
		}



	default:
		send(conn, []byte("invalid command"))
	}
	return nil
}
