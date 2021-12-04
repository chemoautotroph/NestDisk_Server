package server

import (
	"fmt"
	"log"
	"myServer/protocol"
	"myServer/utils"
	"net"
	"strings"
)

type Job interface {
	Do()
}

type JobQueueInfo struct {
	Num  int
	Conn net.Conn
}

func (d *JobQueueInfo) Do() {
	buffer := make([]byte, 1024*1024)
	for{
		tempbuf := make([]byte, 0)
		readerChannel := make(chan []byte, 16)

		n, _ := d.Conn.Read(buffer)
		// close conn at here?
		tempbuf = protocol.Depack(append(tempbuf, buffer[:n]...), readerChannel)
		err := read(d.Conn, readerChannel)
		if err != nil {
			utils.ErrorRecorder(err)
		}
	}
}

func read(conn net.Conn, readerChannel chan []byte) error {
	select {
	case data := <-readerChannel:
		// fmt.Println(string(data))
		connInput := strings.Split(string(data), " ")
		fCommand := connInput[0]
		rCommand := connInput[1:]
		fmt.Printf("first command is: %v\n", fCommand)
		err := runCommand(fCommand, rCommand, conn)
		if err != nil {
			return err
		}
	}
	return nil
}

func send(conn net.Conn, message []byte) {
	//session := strconv.FormatInt(time.Now().Unix(), 10)
	//message = append([]byte(session), message...)
	_, err := conn.Write(protocol.Enpack(message))
	if err != nil {
		log.Fatalln("conn.Write Unexpected Error", err)
	}
}

