package namespaces

import (
	"sync"
	"fmt"
)

var mutex *sync.Mutex = &sync.Mutex{}
var callbacks map[int] func() error = make(map[int] func() error, 20)

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

type NamespaceId int

type Namespaces interface {
	Add(id... NamespaceId) Namespaces
	CloneFlags() int
}

type namespaces map[NamespaceId] bool

func New(ids... NamespaceId) Namespaces {
	ns := namespaces(make(map[NamespaceId] bool, 20))
	ns.Add(ids...)
	return ns
}

func (ns namespaces) Add(ids... NamespaceId) Namespaces {
	for _, id := range ids {
		ns[id] = true
	}

	return ns
}

func (ns namespaces) CloneFlags() int {
	var cloneFlags int
	for id, _ := range ns {
		cloneFlags |= int(id)
	}
	return cloneFlags
}
