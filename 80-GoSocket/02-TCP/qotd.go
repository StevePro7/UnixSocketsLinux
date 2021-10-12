package main

import (
	"fmt"
	"log"
	"net"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	con, err := net.Dial("tcp", "Alienware:17")
	checkErr(err)
	defer func(con net.Conn) {
		err := con.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(con)

	msg := ""
	_, err = con.Write([]byte(msg))
	checkErr(err)

	reply := make([]byte, 1024)
	_, err = con.Read(reply)

	checkErr(err)
	fmt.Println("reply: ", string(reply))
}
