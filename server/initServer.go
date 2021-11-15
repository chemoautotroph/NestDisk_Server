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

func Init() {
	port := config.Config.GetString("port")
	listener, err := net.Listen("tcp", port)
	utils.ErrorRecorder(err)
	pool := initWorkerPool()
	go activatedWorker()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener error %v\n", err)
			continue
		}
		//go handleConn(conn)
		go giveJob(conn, pool)
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

func initWorkerPool() *WorkerPool {
	workerLen := config.Config.GetInt("workerLen")
	p := NewWorkerPool(workerLen)
	p.Run()
	return p
}

func giveJob(conn net.Conn, pool *WorkerPool) {
	sc := &JobQueueInfo{Conn: conn}
	pool.JobQueue <- sc
}

func activatedWorker (){
	for{
		time.Sleep(2 * time.Second)
		fmt.Printf("goroutine: %v\n", runtime.NumGoroutine())
	}
}