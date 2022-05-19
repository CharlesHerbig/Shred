package Shred

import (
	"testing"
)

func Test_root_Shred(t *testing.T) {
	// A file that the user doesnot own.
	err := Shred("rootfile")
	if err != nil {
		t.Error("Got an error during root test", err)
	}
}
