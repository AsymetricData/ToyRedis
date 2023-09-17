package main

import (
	"fmt"
	"os"

	// Uncomment this block to pass the first stage
	"net"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	s, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	_, err = s.Accept()

	if err != nil {
		fmt.Println("Server failed to Accept ", err.Error())
		os.Exit(1)
	}

}
