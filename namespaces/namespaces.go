package namespaces

import (
	"sync"
	"fmt"
)

var cloneFlags int
var mutex *sync.Mutex = &sync.Mutex{}
var callbacks []func() error

func AddCloneFlag(cloneFlag int) {
	mutex.Lock()
	cloneFlags |= cloneFlag
	mutex.Unlock()
}

func CloneFlags() int {
	return cloneFlags
}

func RegisterCallback(callback func() error) {
	mutex.Lock()
	callbacks  = append(callbacks, callback)
	mutex.Unlock()
}

func InNamespaces() error {
	mutex.Lock()
	cbs := callbacks
	mutex.Unlock()
	for _, cb := range cbs {
		if err := cb(); err != nil {
			return fmt.Errorf("Namespace callback %s failed: %s", cb, err.Error())
		}
	}
	return nil
}
