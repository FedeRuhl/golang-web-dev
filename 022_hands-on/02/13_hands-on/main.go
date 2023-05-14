package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
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
	var method string
	var url string
	i := 0

	for scanner.Scan() {
		ln := scanner.Text()

		if ln == "" {
			break
		}

		if i == 0 {
			requestLine := strings.Split(ln, " ")
			method, url = requestLine[0], requestLine[1]
			fmt.Printf("METHOD: %s\r\n", method)
			fmt.Printf("URL: %s\r\n", url)
		} else {
			fmt.Println(ln)
		}

		i++
	}

	body := `<html>
		<head>
			<title>Testing with GO</title>
		</head>
		<body>
			<h1>Hi since GO!</h1>
			<p>Method: %s</p>
			<p>URL: %s</p>
		</body>
	</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	writer := bufio.NewWriter(conn)
	fmt.Fprintf(writer, body, method, url)
	writer.Flush()

	defer conn.Close()
}
