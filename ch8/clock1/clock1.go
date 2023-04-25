package clock1

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("10:15:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func Run() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			return
		}
		handleConn(conn)
	}
}
