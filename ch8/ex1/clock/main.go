package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8000, "port number")
	flag.Parse()

	address := fmt.Sprintf("localhost:%d", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Started listening on port %d \n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format(time.RFC1123+"\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
