package travis

import (
	"testing"
	_ "github.com/nshah/go.subset" // to check that imports are handled
)

func TestMeta(t *testing.T) {
	t.Log("TestMeta works.")
}
