package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"myServer/config"
	"myServer/utils"
	"net"
	"os"
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
		//sendLogo(conn)
		log.Println("conn: ", conn)
		//go handleConn(conn)
		go connHandler(conn, pool)
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

func connHandler(conn net.Conn, pool *WorkerPool) {
	if conn != nil{
		sc := &JobQueueInfo{Conn: conn}
		pool.JobQueue <- sc
	} else {
		log.Panic("invalid socket connection")
	}

}

func activatedWorker (){
	for{
		time.Sleep(2 * time.Second)
		fmt.Printf("goroutine: %v\n", runtime.NumGoroutine())
	}
}

func sendLogo (conn net.Conn){
	file, err := os.Open("./logo.txt")
	if err != nil{
		fmt.Println("fail to open with err ", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		str, err:= reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		send(conn, []byte(str))
	}
}