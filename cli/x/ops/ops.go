package ops

import "os"

// FileExists checks for the existence of the file indicated by filename and returns true if it exists.
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
