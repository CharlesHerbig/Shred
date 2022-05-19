//Package Shred provides a function that overwrites a file with random bytes before deleting it
package Shred

import (
	"crypto/rand"
	"fmt"
	"os"
)

// The Shred function takes one parameter, the file pathname
// It returns an error if it was unable to write or delete the file
// It returns nil if successful
func Shred(pathName string) error {
	// Stat the file for the file length
	fstat, err := os.Stat(pathName)
	if err != nil {
		// File stat failed, return the error
		return fmt.Errorf("in Shred: %w", err)
	}
	var fSize int64 = fstat.Size()
	if fSize == 0 {
		err = os.Remove(pathName)
		if err != nil {
			// Deleting file failed
			return fmt.Errorf("in Shred: %w", err)
		}
		return nil
	}

	// Open the file
	f, err := os.OpenFile(pathName, os.O_WRONLY, 0777)
	if err != nil {
		// File open failed, return the error
		return fmt.Errorf("in Shred: %w", err)
	}
	defer f.Close()

	junkBuf := make([]byte, 1024)
	// Write randdom bytes over the file 3 times
	for i := 0; i < 3; i++ {
		f.Seek(0, 0)
		for fSize = fstat.Size(); fSize > 1024; fSize -= 1024 {
			// Load a buffer with random data
			_, err = rand.Read(junkBuf)
			if err != nil {
				// Getting random bytes failed
				return fmt.Errorf("in Shred: %w", err)
			}
			// Write random bytes to file
			_, err = f.Write(junkBuf)
			if err != nil {
				// Write failed
				return fmt.Errorf("in Shred: %w", err)
			}
		}
		_, err = rand.Read(junkBuf[:fSize])
		if err != nil {
			// Getting random bytes failed
			return fmt.Errorf("in Shred: %w", err)
		}
		_, err = f.Write(junkBuf[:fSize])
		if err != nil {
			// Write failed
			return fmt.Errorf("in Shred: %w", err)
		}
		err = f.Sync()
		if err != nil {
			// Sync to disk failed
			return fmt.Errorf("in Shred: %w", err)
		}
	}
	// Close the file and delete it
	f.Close()
	err = os.Remove(pathName)
	if err != nil {
		// Deleting file failed
		return fmt.Errorf("in Shred: %w", err)
	}
	return nil
}
