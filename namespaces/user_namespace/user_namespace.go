// +build linux
package user_namespace


import (
	"syscall"
)

const Id = syscall.CLONE_NEWUSER
