package clockwall

import (
	"fmt"
	"log"
	"net"
)

func Connect(addr string) net.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func MustRead(src net.Conn, dst chan string) {

}

func Run() {
	conn1 := Connect("localhost:8080")
	//conn2 := Connect("localhost:8081")
	defer conn1.Close()
	//defer conn2.Close()

	out := make([]byte, 10)
	_, err := conn1.Read(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
