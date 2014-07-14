package collecting

import (
	api "github.com/cf-guardian/prototype/libcontainer_api"
	naming "github.com/cf-guardian/prototype/libcontainer_api/naming"
)

// Manage a collection of named containers.
// Each container is represented by its internal name (CName), and associated with its (external) Name by
// the collection.
type Collector interface {

	// Add a Container (identified by cname) to this collection with the given name.
	Add(name api.Name, cname naming.CName) error

	// Remove the given name from this collection.
	Remove(name api.Name) error

	// Return the internal CName associated with the given name from this collection.
	Get(name api.Name) (naming.CName, error)

	// Return the names in this collection.
	Names() []api.Name

}
