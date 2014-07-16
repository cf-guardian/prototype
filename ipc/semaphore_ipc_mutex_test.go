package ipc_test

import (
	"github.com/cf-guardian/prototype/ipc"
	"testing"
)

func TestOpen(t *testing.T) {
	sem, err := ipc.GetSemaphoreIPCMutex("test-sem")
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}

	err = sem.Destroy()
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}
}

