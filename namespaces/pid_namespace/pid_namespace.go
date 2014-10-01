// +build linux
package pid_namespace


import (
	"syscall"
)

const Id = syscall.CLONE_NEWPID
