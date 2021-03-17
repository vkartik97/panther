package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/vkartik97/panther/store"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	message := strings.Split(strings.ToLower(string(buffer[:len(buffer)-2])), " ")
	output := ""
	if message[0] == "get" {
		output = store.Read(message[1])
	} else {
		store.Write(message[1], message[2])
		output = "OK"
	}

	output += "\n"

	conn.Write([]byte(output))

	handleConnection(conn)
}
