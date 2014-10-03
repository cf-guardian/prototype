// +build linux
package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"github.com/cf-guardian/prototype/namespaces"
)

type container struct {
	cmd *exec.Cmd
}

func CreateContainer(ns namespaces.Namespaces, executable string, args... string) (Container, error) {
	err := checkRoot()
	if err != nil {
		return nil, err
	}

	cloneFlags := ns.CloneFlags()

	initArgs := make([]string, 0, len(args)+1)
	initArgs = append(initArgs, fmt.Sprintf("%x", cloneFlags))
	initArgs = append(initArgs, executable)
	initArgs = append(initArgs, args...)

	cmd := exec.Command("init", initArgs...)

	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.Cloneflags = uintptr(cloneFlags)
	cmd.SysProcAttr.Pdeathsig = syscall.SIGKILL

	// DEBUG
//	if output, err := cmd.CombinedOutput(); err != nil {
//		return nil, fmt.Errorf("DEBUG cmd.CombinedOutput error", err.Error())
//	} else {
//		fmt.Println(string(output))
//	}
	// DEBUG

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	go func() {
		data := make([]byte, 1024)
		for {
			n, err := stdOut.Read(data)
			if n > 0 {
				_, err = fmt.Println(string(data[:]))
				if err != nil {
					return
				}
			}
			if err != nil {
				fmt.Printf("Error reading standard output pipe: %s\n", err)
				return
			}
		}
	}()
	go func() {
		data := make([]byte, 1024)
		for {
			n, err := stdErr.Read(data)
			if n > 0 {
				_, err = fmt.Println(string(data[:]))
				if err != nil {
					return
				}
			}
			if err != nil {
				fmt.Printf("Error reading standard error pipe: %s\n", err)
				return
			}
		}
	}()

	err = cmd.Start()

	return &container{cmd}, err
}

func (c *container) Terminate() error {
	return c.cmd.Wait()
}

func checkRoot() error {
	if uid := os.Getuid(); uid != 0 {
		return fmt.Errorf("CreatePidContainer must be run as root. Getuid returned %d", uid)
	}
	return nil
}
