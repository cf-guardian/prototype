package libcontainer_api

type Name string

type Factory interface {

	// Creates a container and associates it with the given name.
	Create(name Name) (Container, error)

	// Creates a container and associates it with a name that is not already in use.
	// Returns the name and the container.
	CreateAndName() (Name, Container, error)

	// Gets the container associated with the given name.
	Get(name Name) (Container, error)

}
