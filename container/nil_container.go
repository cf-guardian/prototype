package container

import "os/exec"

type nil_container struct {
	cmd *exec.Cmd
}

func CreateNilContainer(executable string, args... string) (Container, error) {
	cmd := exec.Command(executable, args...)
	err := cmd.Start()

	return &nil_container{cmd}, err
}

func (c *nil_container) Terminate() error {
	return c.cmd.Wait()
}
