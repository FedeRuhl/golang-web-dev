package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()

		if ln == "" {
			break
		}

		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected.")
}
