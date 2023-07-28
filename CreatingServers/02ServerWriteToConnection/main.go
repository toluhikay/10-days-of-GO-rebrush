package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// listen to a server first
	li, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	// loop through ann continue t listen to connection if it's accepted or not
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, "I am doing pretty fine with y go web deve rebrush stuffs")
		fmt.Fprintf(conn, "%v", "You can't blame for wanting a better life for me")
		fmt.Fprintln(conn, "This is me trying to strengthen my GO skilss and there is no error is tsarting over again and again")

		conn.Close()
	}

	// basics of starting a server
	// 1 you listen
	// listen allows you to accept
	// 2 accept gives you ability to either read to or write to a connect and
	// you can close the connection once you like
}
