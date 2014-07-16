package identity

import (
	api "github.com/cf-guardian/prototype/libcontainer_api"
)

// Generate identifiers suitable for Containers.
type Identifier interface {
	// Return an identifier distinct from all other identifiers produced in this host
	Generate() api.Id

	// Converts a Container identifier into a name suitable for use on the API.
	Name(id api.Id) api.Name
}
