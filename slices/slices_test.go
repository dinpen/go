package slices_test

import (
	"testing"

	"github.com/srcio/go/slices"
)

func TestShuffle(t *testing.T) {
	in := []string{"Hello", "Jay", "World", "Pesco", "Chou", "Ding"}
	slices.Shuffle(in)
	t.Log(in)
}

func TestReverse(t *testing.T) {
	in := []string{"Hello", "Jay", "World", "Pesco", "Chou", "Ding"}
	slices.Reverse(in)
	t.Log(in)
}
