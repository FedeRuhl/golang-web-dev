package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	requestLine := request(conn)
	method := requestLine[0]
	url := requestLine[1]

	// write response
	router(conn, method, url)
}

func request(conn net.Conn) [2]string {
	i := 0
	scanner := bufio.NewScanner(conn)
	var method, url string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			requestLine := strings.Fields(ln)
			method, url = requestLine[0], requestLine[1]
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}

	return [2]string{
		method,
		url,
	}
}

func router(conn net.Conn, method string, url string) {
	switch url {
	case "/contact":
		if method == "GET" {
			respondContact(conn)
		} else {
			respondContact(conn) // another method...
		}
	default:
		if method == "GET" {
			respondHome(conn)
		} else {
			respondHome(conn) // another method...
		}
	}
}

func respondContact(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Contact us</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respondHome(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Home</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
