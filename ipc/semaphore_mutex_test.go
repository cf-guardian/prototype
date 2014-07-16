package ipc_test

import (
	"github.com/cf-guardian/prototype/ipc"
	"testing"
)

func TestGet(t *testing.T) {
	sem, err := ipc.GetSemaphoreMutex("test-sem")
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

func TestLock(t *testing.T) {
	sem, err := ipc.GetSemaphoreMutex("test-sem")
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}

	err = sem.Lock()
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}

	err = sem.Unlock()
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

func TestClose(t *testing.T) {
	sem, err := ipc.GetSemaphoreMutex("test-sem")
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}

	err = sem.Close()
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

func TestTryLock(t *testing.T) {
	sem, err := ipc.GetSemaphoreMutex("test-sem")
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}

	err = sem.TryLock()
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}

	err = sem.Unlock()
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

// TODO: is there some way of policing this and removing the restriction on Unlock?
//func TestUnmatchedUnlock(t *testing.T) {
//	sem, err := ipc.GetSemaphoreMutex("test-sem")
//	if err != nil {
//		t.Errorf("Failed: %s", err)
//		return
//	}
//
//	err = sem.Unlock()
//	if err == nil {
//		t.Error("Unmatched unlock should not have succeeded")
//		return
//	}
//
//	err = sem.Destroy()
//	if err != nil {
//		t.Errorf("Failed: %s", err)
//		return
//	}
//}

