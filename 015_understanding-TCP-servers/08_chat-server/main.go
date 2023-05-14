package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	li, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer li.Close()
	fmt.Println("Listening on port:", PORT)

	conn, err := li.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message == "exit\n" {
			break
		}
		fmt.Println("->: " + message)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		// newmessage := strings.ToUpper(message)
		// conn.Write([]byte(newmessage + "\n"))
		conn.Write([]byte(text + "\n"))
	}
}
