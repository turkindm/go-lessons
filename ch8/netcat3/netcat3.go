package netcat3

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Run() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func(done chan struct{}) {
		io.Copy(os.Stdout, conn)
		fmt.Println("done")
		done <- struct{}{}
	}(done)
	io.Copy(conn, os.Stdin)

	<-done
	tcp := conn.(*net.TCPConn)
	err = tcp.CloseRead()
	if err != nil {
		log.Fatal(err)
	}
	err = tcp.CloseRead()
	if err != nil {
		log.Fatal(err)
	}
}
