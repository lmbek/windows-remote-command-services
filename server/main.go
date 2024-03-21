// server
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientConn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error connecting to client:", err)
		return
	}
	defer clientConn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := scanner.Text()
		fmt.Println("Received command from master client:", command)

		// Forward command to client
		_, err := clientConn.Write([]byte(command + "\n"))
		if err != nil {
			fmt.Println("Error forwarding command to client:", err)
			continue
		}

		// Receive response from client
		clientResponse := make([]byte, 4096) // Assuming response won't exceed 4096 bytes
		n, err := clientConn.Read(clientResponse)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error receiving response from client:", err)
			}
			break
		}

		fmt.Println("Received response from client:", string(clientResponse[:n]))

		// Send response back to master client
		_, err = conn.Write(clientResponse[:n])
		if err != nil {
			fmt.Println("Error sending response to master client:", err)
			continue
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server listening on 127.0.0.1:8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
