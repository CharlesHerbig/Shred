package Shred

import (
	"testing"
)

func Test_empty_Shred(t *testing.T) {
	// An empty file that the user has access to.
	err := Shred("emptyfile")
	if err != nil {
		t.Error("Got an error during empty file test", err)
	}
}
