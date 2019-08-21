package lterator

import (
	"testing"
)

func TestIterator(t *testing.T) {

	names := []string{"tim", "bob", "jet", "jack"}
	nameRepository := NewNameRepository(names)

	iterator := nameRepository.GetIterator()

	for iterator.HasNext() {
		t.Logf("name:%v", iterator.Next() )
	}
}
