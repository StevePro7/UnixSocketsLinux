package main

import (
	"io"
	"log"
	"net"
	"os"
)

const SockAddr = "/tmp/echo.sock"

func echoServer(c net.Conn) {
	address := c.RemoteAddr()
	network := address.Network()
	//log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	log.Printf("Client connected [%s]", network)

	_, err := io.Copy(c, c)
	if err != nil {
		return
	}
	err = c.Close()
	if err != nil {
		return
	}
}

func main() {
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			log.Fatal("close error: ", err)
		}
	}(l)

	for {
		// Accept new connections, dispatching them to echoServer in a goroutine
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error: ", err)
		}

		go echoServer(conn)
	}
}