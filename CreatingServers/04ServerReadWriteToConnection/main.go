package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// reading from a connection using the bufio.Scanner
	li, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

// create a helper function to help read conn with bufio.Scanner
func handle(conn net.Conn) {
	//  i could set a deadline to abort connections if none
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONNECTION TIMED OUT")
	}

	scanner := bufio.NewScanner(conn)

	// check if there is a string to scan using a continuous loop
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "%v\n", ln)
	}
	defer conn.Close()
	fmt.Println("COde reaches here")
}
