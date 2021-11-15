package server

import (
	"fmt"
	"myServer/protocol"
	"net"
)

type Job interface {
	Do()
}

type JobQueueInfo struct {
	Num  int
	Conn net.Conn
}

func (d *JobQueueInfo) Do() {
	tempbuf := make([]byte, 0)
	readerChannel := make(chan []byte, 16)


	buffer := make([]byte, 1024*1024)
	n, _ := d.Conn.Read(buffer)
	tempbuf = protocol.Depack(append(tempbuf, buffer[:n]...), readerChannel)
	reader(readerChannel)
}

func reader(readerChannel chan []byte) {
	select {
	case data := <-readerChannel:
		fmt.Println(string(data))
	}

}
