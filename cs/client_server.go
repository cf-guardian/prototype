package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"os"
	"strings"
	"time"
)

func main() {

	args := os.Args
	if len(args) > 1 && args[1] == "server" {
		server()
	} else {
		fmt.Println("About to start the server")
		cmd := exec.Command(args[0], "server")
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		time.Sleep(100 * time.Millisecond)

		client()

		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
	}
}

func client() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:2345")
	if err != nil {
		log.Fatal(err)
	}

	connection, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	send(connection, "test yyy\n")

	fmt.Println("Result: ", receive(connection))

	send(connection, "exit\n")
}

func server() {
	listener, err := net.Listen("tcp", ":2345")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(connection net.Conn) {
			defer connection.Close()

			req := receive(connection)
			if req == "exit\n" {
				connection.Close()
				os.Exit(0)
			}

			response := fmt.Sprintf("response to %s\n", req)
			send(connection, response)

			fmt.Println("Connection closed on server")
		}(connection)
	}
}

func send(connection net.Conn, request string) {
	if _, err := connection.Write([]byte(request)); err != nil {
		log.Fatal(err)
	}
}

func receive(connection net.Conn) string {
	result := ""
	done := false
	for !done {
		reply := make([]byte, 1024)
		if _, err := connection.Read(reply); err != nil {
			log.Fatal(err)
		}
		rep := string(reply)
		result = result + rep
		if strings.Index(rep, "\n") != -1 {
			done = true
		}
	}
	return result
}
