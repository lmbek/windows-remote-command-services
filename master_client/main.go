// master_client
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func receiveResponse(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server closed the connection.")
				return
			}
			fmt.Println("Error reading response from server:", err)
			return
		}
		fmt.Print("Response from server:", response)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Start a goroutine to continuously receive responses from the server
	go receiveResponse(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command to send to clients: ")
		scanner.Scan()
		command := scanner.Text()

		// Send command to server
		_, err := conn.Write([]byte(command + "\n"))
		if err != nil {
			fmt.Println("Error sending command to server:", err)
			continue
		}
	}
}
