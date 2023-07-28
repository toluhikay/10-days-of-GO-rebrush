package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("I dialied you")
}
