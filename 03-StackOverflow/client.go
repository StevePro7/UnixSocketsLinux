package main

import (
	"io"
	"log"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("Client got: ", string(buf[0:n]))
	}
}

func main() {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(c)

	go reader(c)
	for {
		_, err := c.Write([]byte("hi"))
		if err != nil {
			log.Fatal("write error: ", err)
			break
		}
		time.Sleep(1e0)
	}
}