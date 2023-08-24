package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
		defer conn.Close()
	}
}
