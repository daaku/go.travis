package travis_test

import (
	"github.com/daaku/go.domain" // to check that xtest imports are handled
	"testing"
)

func TestMeta(t *testing.T) {
	dom, _ := domain.TLD("facebook.com")
	if dom != "com" { // just need to use domain
		t.Fatal("Horribly wrong")
	}
	t.Log("TestMeta works.")
}
