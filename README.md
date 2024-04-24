# Go TCP Server

This repository contains a simple TCP server implemented in Go. The server listens for incoming TCP connections and responds with a predefined message. It's a basic example to demonstrate TCP server functionality in Go, useful for learning and prototyping.

## Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.17+ recommended)

## Getting Started

To run the TCP server, follow these steps:
1. Clone the repository or download the code.
```sh
git clone https://github.com/EdAlekseiev/tcp-server.git
cd tcp-server
```
2. Build the server executable.
```sh
go build -o tcpserver ./cmd/tcpserver/main.go
```
3. Start the server. By default, the server listens on port 8085.
```sh
./tcpserver
```

## Connecting to the Server

You can use a TCP client to connect to the server. Here's an example using telnet:
```sh
telnet  localhost  8085
```
