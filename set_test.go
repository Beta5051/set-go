package set

import (
	"testing"
)

func Test_NewSet(t *testing.T) {
	set := NewSet[int]()
	set.Add(1, 3)
	t.Log(set)
}
