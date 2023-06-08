//go:build windows
// +build windows

package secureexec

import (
	"os/exec"

	"github.com/cli/safeexec"
)

// calling exec.Command directly on a windows machine poses a security risk due to
// the current directory being searched first before any directories in the PATH
// variable, meaning you might clone a repo that contains a program called 'git'
// which does something malicious when executed.

// see https://github.com/golang/go/issues/38736 for more context. We'll likely
// be able to just throw out this code and switch to the official solution when it exists.

// I consider this a minor security concern because you're just as vulnerable if
// you call `git status` from the command line directly but no harm in playing it
// safe.

var pathCache = map[string]string{}

func Command(name string, args ...string) *exec.Cmd {
	path := getPath(name)

	return exec.Command(path, args...)
}

func getPath(name string) string {
	if path, ok := pathCache[name]; ok {
		return path
	}

	path, err := safeexec.LookPath(name)
	if err != nil {
		pathCache[name] = name
		return name
	}

	pathCache[name] = path
	return path
}
