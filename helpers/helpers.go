package helpers

import (
	"errors"
	"io/fs"
	"os"
)

// FileExists returns a bool indicating whether a file exists or not. Also
// in the case the input filename is a directory, even if it "exists", the function
// will also return false as it is not a file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}
