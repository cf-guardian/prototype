package identity

import (
	api "github.com/cf-guardian/prototype/libcontainer_api"
)

// Container identifier; should be defined in the base Container interface
type Id string

// Generate identifiers suitable for Containers.
type Identifier interface {
	// Return an identifier distinct from all other identifiers produced in this host
	Generate() Id

	// Converts a Container identifier into a name suitable for use on the API.
	Name(id Id) api.Name
}
