package travis

import (
	_ "github.com/nshah/go.subset" // to check that imports are handled
	"testing"
)

func TestMeta(t *testing.T) {
	t.Log("TestMeta works.")
}
