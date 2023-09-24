package main

import (
	"fmt"
	"io"
	Commands "main/app/Command"
	"main/app/RedisParser"
	"net"
	"os"
	"sync"
)

var db = make(map[string]string)

var mutex sync.Mutex

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	s, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		fmt.Println("Failed to bind to port 12345")
		os.Exit(1)
	}

	defer s.Close()

	for {
		conn, err := s.Accept()
		if err != nil {
			fmt.Println("Server failed to Accept ", err.Error())
			os.Exit(1)
		}

		for {
			buffer := make([]byte, 128)
			_, err = conn.Read(buffer)

			if err != nil {
				fmt.Println("Failed to read data ", err.Error())
			}

			if err == io.EOF {
				fmt.Println("EOF - closing the connection")
				conn.Close()
				break
			}

			commands := RedisParser.ParseBuffer(buffer)

			handleCommand(commands, conn, db)
		}
	}
}

func handleCommand(commands []Commands.Command, conn net.Conn, db map[string]string) {
	for _, command := range commands {
		switch command.Cmd {
		case Commands.PING:
			_, err := conn.Write([]byte("+PONG\r\n"))

			if err != nil {
				fmt.Println(err)
			}
		case Commands.GET:
			mutex.Lock()
			defer mutex.Unlock()
			value, ok := db[*command.Key]
			if !ok {
				println("Not found")
				value = "-ERR"
			} else {
				value = "+" + value
			}
			conn.Write([]byte(string(value + "\r\n")))
		case Commands.SET:
			mutex.Lock()
			defer mutex.Unlock()
			db[*command.Key] = *command.Value
			conn.Write([]byte("+OK\r\n"))
		default:
			fmt.Println("Defaulted")
		}
	}
}
