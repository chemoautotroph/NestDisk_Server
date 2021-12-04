package server

import (
	"fmt"
	"net"
)

func runCommand (fCommand string, rCommand []string, conn net.Conn) error {
	//fmt.Printf("ping to byte: %v\n", []byte("ping"))
	//fmt.Println("fCommand to byte",[]byte(fCommand))

	switch fCommand {
	case "ping":
		send(conn, []byte("pang"))
		fmt.Println("pang")
	case "login":
		if len(rCommand) != 2 {
			send(conn, []byte("invalid Input."))
			break
		}
		message, err, errorMessage := login(rCommand[0], rCommand[1])
		if errorMessage != ""{
			send(conn, []byte(errorMessage))
		} else {
			send(conn, []byte(message))
		}
		if err != nil {
			return err
		}
	}
	return nil
}