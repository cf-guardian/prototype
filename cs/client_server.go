package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"github.com/cf-guardian/prototype/container"
	"strings"
	"time"
	"os/exec"
	"github.com/cf-guardian/prototype/utils"
	"syscall"
	"github.com/cf-guardian/prototype/namespaces"
	"github.com/cf-guardian/prototype/namespaces/ipc_namespace"
	"github.com/cf-guardian/prototype/namespaces/mount_namespace"
//	"github.com/cf-guardian/prototype/namespaces/network_namespace"
	"github.com/cf-guardian/prototype/namespaces/pid_namespace"
//	"github.com/cf-guardian/prototype/namespaces/user_namespace"
	"github.com/cf-guardian/prototype/namespaces/uts_namespace"
)

func main() {

	utils.OptimiseScheduling()

	args := os.Args
	if len(args) > 1 && args[1] == "server" {
		server()
	} else {
		// TODO: add network namespace. This will require veth support.
		// TODO: add user namespace support when Go 1.4 is available.
		ns := namespaces.New(mount_namespace.Id, pid_namespace.Id, uts_namespace.Id, ipc_namespace.Id/*, network_namespace.Id */ /*, user_namespace.Id*/)

		var c container.Container
		var err error
		if c, err = container.CreateContainer(ns, args[0], "server"); err != nil {
			log.Fatal(err)
		}

		time.Sleep(200 * time.Millisecond)

		client()

		if err := c.Terminate(); err != nil {
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
			for {
				req := receive(connection)
				if req == "exit\n" {
					connection.Close()
					os.Exit(0)
				}

				// Gather evidence of being inside a PID namespace
				var evidence string
				evidence = getEvidence()

				response := fmt.Sprintf("response to %s\nEvidence:\n%s\n", strings.Trim(req, "\n"), evidence)
				send(connection, response)
			}
		}(connection)
	}
}

const defaultMountFlags = syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV

func getEvidence() string {
//	// Remount /proc
//	if err := syscall.Unmount("/proc", syscall.MNT_DETACH); err != nil {
//		return fmt.Sprintf("syscall.Unmount error", err.Error())
//	}
//
//	if err := syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), ""); err != nil {
//		return fmt.Sprintf("syscall.Mount error", err.Error())
//	}

	cmd := exec.Command("/bin/ps", "-uf")

	bytes, err := cmd.Output()
	if (err != nil) {
		return fmt.Sprintf("cmd.Output error", err.Error())
	}

	cmd.Wait()
	return string(bytes)
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
		n, err := connection.Read(reply)
		if err != nil {
			log.Fatal(err)
		}
		rep := string(reply[:n])
		result = result + rep
		if strings.Index(rep, "\n") != -1 {
			done = true
		}
	}
	return result
}
