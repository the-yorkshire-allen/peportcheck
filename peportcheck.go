package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: peportcheck [port number] [port number n]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Input ports are missing.")
		os.Exit(1)
	}

	for port := range args {
		go checkPort(args[port])
	}

	for {
		time.Sleep(1000)
	}
}

func checkPort(port string) {
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid port %q: %s", port, err)
		return
	}

	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s", port, err)
		return
	}

	fmt.Printf("TCP Port %q is available\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Printf("Connection accepted from %s on Port %q\n", conn.RemoteAddr(), port)

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Process and use the data (here, we'll just print it)
		fmt.Printf("Received: %s\n", buffer[:n])
	}
}
