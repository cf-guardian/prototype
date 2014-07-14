package collecting

import (
	"github.com/cf-guardian/prototype/libcontainer_api"
	"github.com/cf-guardian/prototype/libcontainer_api/naming"
)

type Collector interface {

	Add(name libcontainer_api.Name, cname naming.CName) error

	Remove(name libcontainer_api.Name) error

	Get(name libcontainer_api.Name) (naming.CName, error)

	Names() []libcontainer_api.Name

}
