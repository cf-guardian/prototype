package libcontainer_api

type Id string

type Container interface {
	Id() Id
}
