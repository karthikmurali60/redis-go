package main

import (
	"fmt"
	"net"
	"os"
	"log"
)

func main() {
	fmt.Println("Logs from your program will appear here!")
	
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handleConnection(connection)
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection: ", err.Error())
			os.Exit(1)
		}
	
		log.Println("Received the following data: ", string(buffer[:n]))
	
		message := []byte("+PONG\r\n")
		_, err = connection.Write(message)
		if err != nil {
			fmt.Println("Error writing to connection: ", err.Error())
			os.Exit(1)
		}
	}
}
