package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		s := strings.Split(arg, "=")
		name, address := s[0], s[1]
		fmt.Println(address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go readTime(name, conn)
		// go mustCopy(os.Stdout, conn)
	}
	for {
		time.Sleep(time.Second)
	}
}

func readTime(name string, src io.Reader) {
	s := bufio.NewScanner(src)
	for s.Scan() {
		fmt.Println(name + " " + s.Text())
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
