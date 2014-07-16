package mutex_test

import (
	"github.com/cf-guardian/prototype/mutex"
	"testing"
)

func TestOpen(t *testing.T) {
	sem, err := mutex.GetSemaphoreMutex("test-sem")
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

