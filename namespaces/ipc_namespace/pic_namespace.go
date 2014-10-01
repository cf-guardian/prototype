// +build linux
package ipc_namespace


import (
	"syscall"
)

const Id = syscall.CLONE_NEWIPC
