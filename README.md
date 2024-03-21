# Windows Remote Command Program

This project is a remote command execution system written in Go. It is designed for controlling multiple PCs simultaneously by sending commands from a master client, which forwards the commands to the server, and the server dispatches the commands to multiple clients. The responses from each client are sent back to the master client, giving you real-time feedback and control over multiple systems.

> ⚠️ **Please Note!** This repository contains code that I made for fun. Keep in mind that this project contains code that can execute windows cmd commands from a master client to many clients through a server (if the IP addresses are modified in the code). Remember, this project is for experimentation and should be utilized for ethical and legal purposes. Misuse of this tool is strictly not endorsed by the author.

## Installation and Running the Services

### Prerequisites
1. Install Go - You need to have Go installed on your system. You can download and install Go from the [official Go website](https://golang.org/dl/). Ensure that you have version `1.22.1` or higher.

### Running the Application

1. **Download the code**
    - Clone the repository to your local machine using `git clone https://github.com/lmbek/windows-remote-command-services.git`.
2. **Navigate to each component's directory**
    - Once you have cloned the repository, navigate into each component's directory (i.e., Client, Master Client, and Server).
3. **Build the applications**
    - While inside each directory, build the Go application using the command:
      ```
      go build main.go
      ```
    - This command will create an executable file in the same directory.
4. **Run the applications**
    - Start each service by issuing the command:
      ```
      go run main.go
      ```
      Remember to run each service in a separate terminal window for them to run concurrently and interact with each other effectively.

> **Note**: For the system to work effectively, the server must be running when the master client sends a command, and at least one client must be active to receive commands from the server.

## Components

This project consists of the following parts:

1. **Client**: A client listens on `127.0.0.1:8081` for incoming TCP connections. When a connection is established, it awaits and executes the commands sent by the server. The command execution is handled on a new goroutine, allowing for concurrent processing on multiple connections.

2. **Master Client**: It establishes a TCP connection to the server at `127.0.0.1:8080`. This part listens for command input from the user through the standard input. Any input command will be sent to the server and the response from the server is displayed to the user.

3. **Server**: The server listens on `127.0.0.1:8080` for incoming TCP connections from the master client, then establishes a TCP connection with the client. The server dispatches the commands received from the master client to the client and handles any responses.

The ip address should be edited to match the specific network

## Running the Services

- Start each service by issuing the `go run main.go` command in the respective component's directory.

> **Note** : Remember to run each service in a separate terminal window as they need to be running concurrently for the system to work effectively.

## Future Plans

The following features are on the roadmap:

- User authentication for security enhancement.
- Encryption of the command and response messages for secure communication.
- Support for more clients and parallel execution.
- Improvements in error handling and feedback.
- Integration with other systems or platforms.
- Adding an interactive user interface for better operation control.

## Contribution

Any contributions from the community are welcome. Please feel free to submit any bugs, requests, or patches.

**Disclaimer**: This project should be used for legal and ethical purposes only.

## License

This project is licensed under the MIT License - see the LICENSE file for details.