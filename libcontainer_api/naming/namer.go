package naming

// Container names; should be defined in the base Container interface
type CName string

// Generate names suitable for Containers.
type Namer interface {
	// Return a name distinct from all other names produced in this host
	Generate() CName
}
