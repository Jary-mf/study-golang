package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type Value interface {
	String() string
	Set(string) error
}

func WebColck_main() {
	var port = flag.Int("port", 8000, "the port you want to connect")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
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
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
