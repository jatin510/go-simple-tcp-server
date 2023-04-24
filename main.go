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
		log.Println("waiting for connection...")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connection established")
		go do(conn)
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
	log.Println("processing the request...")
	time.Sleep(8 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))

}
