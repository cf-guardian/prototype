package namespaces

import (
	"sync"
	"fmt"
)

var cloneFlags int
var mutex *sync.Mutex = &sync.Mutex{}
var callbacks map[int] func() error = make(map[int] func() error, 20)

func AddCloneFlag(cloneFlag int) {
	mutex.Lock()
	cloneFlags |= cloneFlag
	mutex.Unlock()
}

func CloneFlags() int {
	return cloneFlags
}

func RegisterCallback(cloneFlag int, callback func() error) {
	mutex.Lock()
	callbacks[cloneFlag] = callback
	mutex.Unlock()
}

func InNamespaces(cloneFlags int) error {
	mutex.Lock()
	cbs := callbacks
	mutex.Unlock()
	for cloneFlag, cb := range cbs {
		if cloneFlags&cloneFlag != 0 {
			if err := cb(); err != nil {
				return fmt.Errorf("Namespace callback %s failed: %s", cb, err.Error())
			}
		}
	}
	return nil
}
