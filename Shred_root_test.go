package Shred

import (
	"testing"
)

func Test_root_Shred(t *testing.T) {
	// A typical file that the user has access to.
	err := Shred("rootfile")
	if err != nil {
		t.Error("Got an error during test", err)
	}
}
