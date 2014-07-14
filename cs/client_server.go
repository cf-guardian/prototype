package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"os"
	"strings"
	"time"
)

type netAccepter struct {
	listener net.Listener
}

func (ca *netAccepter) Accept() (io.ReadWriter, error) {
	return ca.listener.Accept()
}

func toAccepter(listener net.Listener) Accepter {
	return &netAccepter{listener}
}

func main() {
	args := os.Args
	if len(args) > 1 && args[1] == "server" {
		listener, err := net.Listen("tcp", ":2345")
		if err != nil {
			log.Fatal(err)
		}
		defer listener.Close()

		server(toAccepter(listener), func(connection io.ReadWriter) (bool, int) {
				for {
					req := receive(connection)
					if req == "exit\n" {
						return true, 0
					}
					response := fmt.Sprintf("response to %s\n", strings.Trim(req, "\n"))
					send(connection, response)
				}
			})
	} else {
		cmd := exec.Command(args[0], "server")

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		time.Sleep(100 * time.Millisecond)

		client()

		if err := cmd.Wait(); err != nil {
			log.Fatalf("cmd.Wait failed: %v", err)
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

func send(connection io.Writer, request string) {
	if _, err := connection.Write([]byte(request)); err != nil {
		log.Fatal(err)
	}
}

func receive(connection io.Reader) string {
	result := ""
	done := false
	for !done {
		reply := make([]byte, 1024)
		n, err := connection.Read(reply)
		if err != nil {
			log.Fatal(err)
		}
		rep := string(reply[:n])
		result = result+rep
		if strings.Index(rep, "\n") != -1 {
			done = true
		}
	}
	return result
}

type Accepter interface {
	Accept() (io.ReadWriter, error)
}

func server(listener Accepter, processor func(connection io.ReadWriter) (bool, int)) {
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn io.ReadWriter) {
			exit, exitStatus := processor(conn)
			conn.Close()
			if exit {
				os.Exit(exitStatus)
			}
		}(connection)
	}
}
