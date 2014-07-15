package identity

import (
	"strings"
	"testing"
)

// Test the SimpleIdentifier implementation of Identifier

func TestSICreate(t *testing.T) {
	ider := CreateSimpleIdentifier()
	id := ider.Generate()
	if sid, snm := string(id), string(ider.Name(id)); sid != snm {
		t.Errorf("id not converted verbatim; id=%q, Name(id)=%q", sid, snm)
	}
}

const (
	LARGEISH_NUMBER int = 1000
)

func TestSIForm(t *testing.T) {
	ider := CreateSimpleIdentifier()

	id := ider.Generate()

	sid := string(id)
	if len(sid) != idLength {
		t.Errorf("Identifiers should be exactly %d characters long, was %q.", idLength, sid)
	}
	if !strings.HasPrefix(sid, idPrefix) {
		t.Errorf("Identifiers should begin with %q.", idPrefix)
	}

	digsId := strings.TrimPrefix(sid, idPrefix)
	if str := strings.Trim(digsId, "0123456789"); "" != str {
		t.Errorf("Identifier contains non-digits after prefix: %q.", str)
	}
}

func TestSIMulti(t *testing.T) {
	var genArr [LARGEISH_NUMBER]Id

	ider := CreateSimpleIdentifier()
	for i := 0; i < LARGEISH_NUMBER; i++ {
		genArr[i] = ider.Generate()
	}

	// Now check they are all different
	set := make(map[Id]struct{})
	for i := 0; i < LARGEISH_NUMBER; i++ {
		set[genArr[i]] = struct{}{}
	}

	if l := len(set); l < LARGEISH_NUMBER {
		t.Errorf("Non-unique identifiers: %d distinct ids returned in first %d.", l, LARGEISH_NUMBER)
	}
}

func TestSIClash(t *testing.T) {
	var genArr [LARGEISH_NUMBER]Id

	ider1 := CreateSimpleIdentifier()
	ider2 := CreateSimpleIdentifier()

	for i := 0; i < LARGEISH_NUMBER; i++ {
		genArr[i] = ider1.Generate()
	}

	set := make(map[Id]struct{})
	for i := 0; i < LARGEISH_NUMBER; i++ {
		set[genArr[i]] = struct{}{}
	}

	// Now check the ones from ider2 don't clash
	for i := 0; i < LARGEISH_NUMBER; i++ {
		gid := ider2.Generate()
		if _, has := set[gid]; has {
			t.Errorf("Identifiers clashed: %q generated before!", gid)
		}
	}

	if l := len(set); l < 2*LARGEISH_NUMBER {
	}
}
