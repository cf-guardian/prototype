// +build linux
package pid_namespace


import (
	"github.com/cf-guardian/prototype/namespaces"
	"syscall"
)

func init() {
	namespaces.AddCloneFlag(syscall.CLONE_NEWPID)
}
