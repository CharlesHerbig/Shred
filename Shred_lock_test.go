package Shred

import (
	"testing"
)

func Test_lock_Shred(t *testing.T) {
	// A file where permission bits deny the user access.
	err := Shred("lockfile")
	if err != nil {
		t.Error("Got an error during locked test", err)
	}
}
