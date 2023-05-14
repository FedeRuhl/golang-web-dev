package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"text/template"
)

type request struct {
	Method string
	Url    string
}

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
	r := request{}
	i := 0

	for scanner.Scan() {
		ln := scanner.Text()

		if ln == "" {
			break
		}

		if i == 0 {
			requestLine := strings.Split(ln, " ")
			r.Method, r.Url = requestLine[0], requestLine[1]
			fmt.Printf("METHOD: %s\r\n", r.Method)
			fmt.Printf("URL: %s\r\n", r.Url)
		} else {
			fmt.Println(ln)
		}

		i++
	}

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", 1000)
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	err = tpl.ExecuteTemplate(conn, "tpl.gohtml", r)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
}
