package collecting

import (
	api "github.com/cf-guardian/prototype/libcontainer_api"
	identity "github.com/cf-guardian/prototype/libcontainer_api/identity"
)

// Manage a collection of named containers.
// Each container is represented by its identity (Id), and associated with its (external) name by
// the collection.
type Collector interface {

	// Add a Container (identified by id) to this collection with the given name.
	Add(name api.Name, id identity.Id) error

	// Remove the given name from this collection.
	Remove(name api.Name) error

	// Return the internal Id associated with the given name from this collection.
	Get(name api.Name) (identity.Id, error)

	// Return the (external) names in this collection.
	Names() []api.Name

}

var theCollector *collector

// Returns a Collector, after either creating one or finding an existing one.
func GetCollector() Collector {
	return theCollector
}

type collector struct {

}

func (coll *collector) Add(name api.Name, id identity.Id) error {
	panic("unimplemented")
}

func (coll *collector) Remove(name api.Name) error {
	panic("unimplemented")
}

func (coll *collector) Get(name api.Name) (identity.Id, error) {
	panic("unimplemented")
}

func (coll *collector) Names() []api.Name {
	panic("unimplemented")
}
