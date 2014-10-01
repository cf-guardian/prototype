package main

import (
	"log"
	"os"
	"github.com/cf-guardian/prototype/namespaces"
	"github.com/cf-guardian/prototype/utils"
	"os/exec"
	_ "github.com/cf-guardian/prototype/namespaces/mount_namespace"
	_ "github.com/cf-guardian/prototype/namespaces/pid_namespace"
	"strconv"
)

// The init process is the process root of a container.
//
// Arguments:
// 0 - the init process path
// 1 - namespace flags
// 2 - the target process path
// 3 - target process arguments
func main() {

	utils.OptimiseScheduling()

	args := os.Args
	if len(args) >= 3 {
		cloneFlags, _ := strconv.ParseInt(args[1], 16, 0)
		if err := namespaces.InNamespaces(int(cloneFlags)); err != nil {
			log.Fatalf("Init process namespace initialisation failed: %s\n", err)
		}
		cmd := exec.Command(args[2], args[3:]...)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Target process failed to run: %s\n", err)
		}
	} else {
		log.Fatalf("Insufficient init process arguments: %v\n", args)
	}
}
