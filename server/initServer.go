package server

import (
	"fmt"
	"goCloud/config"
	"goCloud/utils"
	"log"
	"net"
)

func init(){
	port := config.Config.GetString("port")
	listener, err := net.Listen("tcp",port)
	utils.ErrorRecorder(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener error %v\n",err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn){
	// defer conn.Close()
	for{
		buf := make([]byte,512)
		n, err := conn.Read(buf)
		if err != nil{
			log.Fatalln("error reading ",err)
			return
		}
		fmt.Printf("Receive data: %v\n", string(buf[:n]))
	}

}

