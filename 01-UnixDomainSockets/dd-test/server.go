package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

const SockAddr = "/tmp/rpc.sock"

type Greeter struct {
}

func (g Greeter) Greet(name *string, reply *string) error {
	log.Printf("Input: name [%s] reply [%s]", *name, *reply)
	*reply = "Hello, " + *name
	return nil
}

func main() {
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	greeter := new(Greeter)
	err := rpc.Register(greeter)
	if err != nil {
		log.Fatal(err)
		return
	}
	rpc.HandleHTTP()
	l, e := net.Listen("unix", SockAddr)
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	fmt.Println("Serving...")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}