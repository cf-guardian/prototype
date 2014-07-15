package identity

import (
	rand "math/rand"
	api "github.com/cf-guardian/prototype/libcontainer_api"
	"time"
	"fmt"
)

type simpleIdentifier struct {
	seed int64
}

// implementation parameters
const (
	idPrefix string = "SI"
	idLength int = 22
	idFormat string = "SI%020d"	// consistent with idPrefix and idLength
)

func CreateSimpleIdentifier() Identifier {
	rand.Seed(time.Now().UTC().UnixNano())

	return &simpleIdentifier{seed: rand.Int63()}
}

func (this *simpleIdentifier) Generate() Id {
	this.seed++
	return Id(fmt.Sprintf(idFormat, this.seed))
}

func (*simpleIdentifier) Name(id Id) api.Name {
	return api.Name(id)
}
