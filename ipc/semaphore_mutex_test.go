package ipc_test

import (
	"github.com/cf-guardian/prototype/ipc"
	"testing"
)

func TestGet(t *testing.T) {
	sem, err := ipc.GetSemaphoreMutex("test-sem")
	if err != nil {
		t.Fatalf("Failed GetSemaphoreMutex: %s", err)
	}

	err = sem.Destroy()
	if err != nil {
		t.Fatalf("Failed Destroy: %s", err)
	}
}

func TestLock(t *testing.T) {
	sem, err := ipc.GetSemaphoreMutex("test-sem")
	if err != nil {
		t.Fatalf("Failed GetSemaphoreMutex: %s", err)
	}

	err = sem.Lock()
	if err != nil {
		t.Fatalf("Failed Lock: %s", err)
	}

	err = sem.Unlock()
	if err != nil {
		t.Fatalf("Failed Unlock: %s", err)
	}

	err = sem.Destroy()
	if err != nil {
		t.Fatalf("Failed Destroy: %s", err)
	}
}

func TestClose(t *testing.T) {
	sem, err := ipc.GetSemaphoreMutex("test-sem")
	if err != nil {
		t.Fatalf("Failed GetSemaphoreMutex: %s", err)
	}

	err = sem.Close()
	if err != nil {
		t.Fatalf("Failed Close: %s", err)
	}

	err = sem.Destroy()
	if err != nil {
		t.Fatalf("Failed Destroy: %s", err)
	}
}

func TestTryLock(t *testing.T) {
	sem, err := ipc.GetSemaphoreMutex("test-sem")
	if err != nil {
		t.Fatalf("Failed GetSemaphoreMutex: %s", err)
	}

	err = sem.TryLock()
	if err != nil {
		t.Fatalf("Failed TryLock: %s", err)
	}

	err = sem.Unlock()
	if err != nil {
		t.Fatalf("Failed Unlock: %s", err)
	}

	err = sem.Destroy()
	if err != nil {
		t.Fatalf("Failed Destroy: %s", err)
	}
}

// TODO: is there some way of policing this and removing the restriction on Unlock?
//func TestUnmatchedUnlock(t *testing.T) {
//	sem, err := ipc.GetSemaphoreMutex("test-sem")
//	if err != nil {
//		t.Fatalf("Failed GetSemaphoreMutex: %s", err)
//	}
//
//	err = sem.Unlock()
//	if err == nil {
//		t.Fatalf("Unmatched unlock should not have succeeded")
//	}
//
//	err = sem.Destroy()
//	if err != nil {
//		t.Fatalf("Failed Destroy: %s", err)
//	}
//}
