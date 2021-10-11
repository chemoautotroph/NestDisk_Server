package server

import (
	"fmt"
	"log"
	"myServer/config"
	"myServer/utils"
	"net"
	"runtime"
	"time"
)

func InitServer() {
	port := config.Config.GetString("port")
	listener, err := net.Listen("tcp", port)
	utils.ErrorRecorder(err)
	workerLen := config.Config.GetInt("workerNumb")
	p := NewWorkerPool(workerLen)
	p.Run()

	go func () {
		for { // 阻塞主程序结束
			fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
			time.Sleep(2 * time.Second)
		}
	} ()

	for i:=0; ; i++{
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener error %v\n", err)
			continue
		}
		sc := &Jobs{Num: i, Conn: conn}
		// log.Printf("给出工作：%v\n", i)
		p.JobQueue <- sc
		// handleConnWithWorkerPool(conn, p)
		// go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	dataBlockSize := 1024 * 100 // 100kb
	err := conn.SetDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		return
	}
	for {
		buf := make([]byte, dataBlockSize)
		n, err := conn.Read(buf)

		if err != nil {
			log.Fatalln("error reading ", err)
			return
		}
		fmt.Printf("Receive data: %v\n", string(buf[:n]))
	}

}

func handleConnWithWorkerPool(conn net.Conn, p *WorkerPool) {
	go func() {
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				log.Printf("err :%v", err)
			}
		}(conn)
		for {
			sc := &Jobs{Num: 1, Conn: conn}
			// log.Printf("给出工作：%v\n", 1)
			p.JobQueue <- sc
		}
	}()
	for { // 阻塞主程序结束
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
