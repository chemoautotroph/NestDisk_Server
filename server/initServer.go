package server

import (
	"fmt"
	"log"
	"myServer/config"
	"myServer/utils"
	"net"
	"time"
)

func Init(){
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
	dataBlockSize := 1024 * 100 // 100kb
	err := conn.SetDeadline(time.Now().Add(time.Second*10))
	if err != nil {
		return
	}
	for{
		buf := make([]byte,dataBlockSize)
		n, err := conn.Read(buf)

		if err != nil{
			log.Fatalln("error reading ",err)
			return
		}
		fmt.Printf("Receive data: %v\n", string(buf[:n]))
	}

}

