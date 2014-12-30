package statement

import (
	"testing"

	"github.com/zimmski/go-mutesting/test"
)

func TestMutatorRemoveStatement(t *testing.T) {
	test.Mutator(
		t,
		NewMutatorRemoveStatement(),
		"../../testdata/statement/remove.go",
		9,
	)
}