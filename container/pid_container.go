// +build linux
package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type pid_container struct {
	cmd *exec.Cmd
}

func CreatePidContainer(executable string, args... string) (Container, error) {
	err := checkRoot()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(executable, args...)

	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.Cloneflags = uintptr(syscall.CLONE_NEWPID|syscall.CLONE_NEWNS)
	cmd.SysProcAttr.Pdeathsig = syscall.SIGKILL

	err = cmd.Start()

	return &nil_container{cmd}, err
}

func (c *pid_container) Terminate() error {
	return c.cmd.Wait()
}

func checkRoot() error {
	if uid := os.Getuid(); uid != 0 {
		return fmt.Errorf("CreatePidContainer must be run as root. Getuid returned %d", uid)
	}
	return nil
}
