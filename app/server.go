package main

import (
	"fmt"
	"os"
	"strings"

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

	defer s.Close()

	for {
		conn, err := s.Accept()
		if err != nil {
			fmt.Println("Server failed to Accept ", err.Error())
			os.Exit(1)
		}

		defer conn.Close()

		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)

		if err != nil {
			fmt.Println("Failed to read data ", err.Error())
		}

		str := strings.Split(string(buffer), "\n")

		fmt.Println(str)

		conn.Write([]byte("+PONG\r\n"))

		/* 		To test your implementation using the official Redis CLI, you can start your server using ./spawn_redis_server.sh and then run echo -e "ping\nping" | redis-cli from your terminal. This will send two PING commands using the same connection. */

	}

}
