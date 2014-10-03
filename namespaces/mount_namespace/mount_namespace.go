// +build linux
package mount_namespace

import (
	"fmt"
	"github.com/cf-guardian/prototype/namespaces"
	"syscall"
)

const Id = syscall.CLONE_NEWNS

const defaultMountFlags = syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV

func init() {
	namespaces.RegisterCallback(syscall.CLONE_NEWNS, initialise)
}

func initialise() error {
	// Remount /proc
	if err := syscall.Unmount("/proc", syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("syscall.Unmount error: %s", err.Error())
	}

	if err := syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), ""); err != nil {
		return fmt.Errorf("syscall.Mount error: %s", err.Error())
	}

	return nil
}
