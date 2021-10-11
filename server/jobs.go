package server

import (
	"log"
	"myServer/Service"
	"net"
	"runtime"
	"time"
)

type Job interface {
	Do()
}

type Jobs struct {
	Num  int
	Type string
	Data []byte
	Conn net.Conn
}

func (j *Jobs) Do() {
	buffer := make([]byte, 1000)
	n, _ := j.Conn.Read(buffer)
	log.Printf("goroutine数量: %v, 已给出的任务：%v, 数据大小: %v bytes, 数据内容: %v\n", runtime.NumGoroutine(),j.Num, n, string(buffer[:n]))
	user := Service.SetUserInfo("insert1", "123456", time.Now())
	user.Insert()
	user.Login()
	// log.Printf("Receive data: %v\n", string(buffer[:n]))

}

func (j *Jobs) identify(){

}
