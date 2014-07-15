package mutex_test

import (
	"github.com/cf-guardian/prototype/libcontainer_api/mutex"
	"testing"
)

func TestOpen(t *testing.T) {
	sem, err := mutex.SemOpen("test-sem")
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

