package server

import (
	"fmt"
	"log"
	"myServer/config"
	"myServer/utils"
	"net"
	"time"
)

func InitServer(){
	port := config.Config.GetString("port")
	listener, err := net.Listen("tcp",port)
	utils.ErrorRecorder(err)
	workerLen := config.Config.GetInt("workerNumb")
	p := NewWorkerPool(workerLen)
	p.Run()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener error %v\n",err)
			continue
		}
		handleConnWithWorkerPool(conn, p)
		log.Println(conn)
		// go handleConn(conn)
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

func handleConnWithWorkerPool(conn net.Conn, p *WorkerPool){
	go func() {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("err :%v",err)
		}
	}(conn)
		for i := 1; ; i++{
			sc := &Jobs{Num:  i, Conn: conn}
			p.JobQueue <- sc
		}
	}()
}
