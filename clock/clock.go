package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // For example, abort connection
			continue
		}

		handleConn(conn) // Handle one connect per time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format(time.RFC1123+"\n"))
		if err != nil {
			return // For example, user disconnec
		}

		time.Sleep(1 * time.Second)
	}
}
