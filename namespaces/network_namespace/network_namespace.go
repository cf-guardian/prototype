// +build linux
package network_namespace


import (
	"syscall"
)

const Id = syscall.CLONE_NEWNET
