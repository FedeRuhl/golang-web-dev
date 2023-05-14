package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"text/template"
)

type request struct {
	Method string
	Url    string
	Body   string
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
	defer conn.Close()

	buf := make([]byte, 1024) // default buffer size?
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// request := string(buf[:n])

	requestLine := string(buf[0:strings.Index(string(buf), "\n")])
	parts := strings.Fields(requestLine)
	method, url := parts[0], parts[1]

	request := request{
		Method: method,
		Url:    url,
		Body:   "",
	}

	if url == "/apply" && method == "GET" {
		handleGetApply(conn, request)
	} else if url == "/apply" && method == "POST" {
		bodyStart := "\r\n\r\n"
		bodyIndex := -1

		for i := 0; i < len(buf); i++ {
			if buf[i] == '\r' && i+3 < len(buf) && string(buf[i:i+4]) == bodyStart {
				bodyIndex = i + 4
				break
			}
		}

		if bodyIndex != -1 {
			body := string(buf[bodyIndex:n])
			nameStartIndex := strings.LastIndex(body, "=")
			request.Body = body[nameStartIndex+1:]
		}
		handlePostApply(conn, request)
	} else {
		handleGetIndex(conn, request)
	}
}

func handleGetIndex(conn net.Conn, r request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", 1024)
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	err = tpl.ExecuteTemplate(conn, "index.gohtml", r)
	if err != nil {
		fmt.Println(err)
	}
}

func handleGetApply(conn net.Conn, r request) {
	tpl, err := template.ParseFiles("form.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", 1024)
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	err = tpl.ExecuteTemplate(conn, "form.gohtml", r)
	if err != nil {
		fmt.Println(err)
	}
}

func handlePostApply(conn net.Conn, r request) {
	tpl, err := template.ParseFiles("form.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", 1024)
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")

	err = tpl.ExecuteTemplate(conn, "form.gohtml", r)
	if err != nil {
		fmt.Println(err)
	}
}
