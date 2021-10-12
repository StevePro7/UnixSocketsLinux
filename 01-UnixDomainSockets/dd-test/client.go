package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("unix", "/tmp/rpc.sock")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	name := "StevePro"
	var reply string
	err = client.Call("Greeter.Greet", &name, &reply)
	if err != nil {
		log.Fatal("greeter error: ", err)
	}
	fmt.Printf("Got '%s'\n", reply)
}