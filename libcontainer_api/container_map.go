package libcontainer_api

import (
	"errors"
	"sync"
	"github.com/cf-guardian/prototype/libcontainer_api/identity"
)

var (
	ErrAlreadyExists error = errors.New("name already associated with a container")
)

func GetLocalFactory() Factory {
	return &containerMap{identifier: identity.CreateSimpleIdentifier()}
}

type containerMap struct {
	mutex sync.Mutex
	container map[Name]Container
	identifier identity.Identifier
}

func (cm *containerMap) Create(name Name) (Container, error) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	_, present := cm.container[name]
	if present {
		return nil, ErrAlreadyExists
	}

	return createWithName(name)
}

// Must be called with the mutex locked.
func (cm *containerMap) createWithName(name Name) (Container, error) {
	if c, err := create(); err == nil {
		cm.container[name] = c
		return c, nil
	} else {
		return nil, err
	}
}

func (cm *containerMap) CreateAndName() (Name, Container, error) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	present := true
	for present {
		id := cm.identifier.Generate()
		name := cm.identifier.Name(id)
		_, present = cm.container[name]
	}


}

func (cm *containerMap) Get(name Name) (Container, error) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

}

func create() (Container, error) {
	return 42, nil
}
