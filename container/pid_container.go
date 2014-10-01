// +build linux
package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"github.com/cf-guardian/prototype/namespaces"
	_ "github.com/cf-guardian/prototype/namespaces/mount_namespace"
	_ "github.com/cf-guardian/prototype/namespaces/pid_namespace"
)

type pid_container struct {
	cmd *exec.Cmd
}

func CreateContainer(executable string, args... string) (Container, error) {
	err := checkRoot()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(executable, args...)

	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.Cloneflags = uintptr(namespaces.CloneFlags())
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
