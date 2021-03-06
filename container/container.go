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

	initArgs := make([]string, 0, len(args) + 1)
	initArgs = append(initArgs, fmt.Sprintf("%x", cloneFlags))
	initArgs = append(initArgs, executable)
	initArgs = append(initArgs, args...)

	cmd := exec.Command("init", initArgs...)

	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.Cloneflags = uintptr(cloneFlags)
	cmd.SysProcAttr.Pdeathsig = syscall.SIGKILL

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
