package gopath

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// Get returns the user's GOPATH, compatible with Go 1.8+. That is,
// it returns the value of the environment variable GOPATH, and if
// that is not set, the directory named "go" in the user's home
// directory.
func Get() (string, error) {
	p := os.Getenv("GOPATH")
	if p != "" {
		return p, nil
	}
	return homedir.Expand("~/go")
}

// Join returns the user's GOPATH plus one or more path elements
// joined with the GOPATH. This works just like filepath.Join.
func Join(elem ...string) (string, error) {
	p, err := Get()
	if err != nil {
		return "", nil
	}

	elem = append([]string{p}, elem...)

	return filepath.Join(elem...), nil
}
