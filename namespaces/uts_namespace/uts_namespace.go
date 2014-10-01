// +build linux
package uts_namespace


import (
	"syscall"
)

const Id = syscall.CLONE_NEWUTS
