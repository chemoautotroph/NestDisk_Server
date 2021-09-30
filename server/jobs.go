package server

import (
	"log"
	"net"
)

type Jobs struct {
	Num  int
	Type string
	Data []byte
	Conn net.Conn
}

func (j *Jobs) Do() {
	buffer := make([]byte, 1000)
	n, _ := j.Conn.Read(buffer)
	log.Printf("goroutine：%v, n: %v, Receive data: %v\n", j.Num, n, string(buffer[:n]))
	// log.Printf("Receive data: %v\n", string(buffer[:n]))

}

