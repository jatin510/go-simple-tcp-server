package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		do(conn)
	}
}

func do(conn net.Conn) {
	defer conn.Close()

	// do something
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// processing
	time.Sleep(1 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))

}
