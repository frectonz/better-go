// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"errors"
	"internal/bytealg"
	"internal/itoa"
	_ "unsafe" // for go:linkname
)

// random number source provided by runtime.
// We generate random temporary file names so that there's a good
// chance the file doesn't exist yet - keeps the number of tries in
// TempFile to a minimum.
//go:linkname runtime_rand runtime.rand
func runtime_rand() uint64

func nextRandom() string {
	return itoa.Uitoa(uint(uint32(runtime_rand())))
}

// CreateTemp creates a new temporary file in the directory dir,
// opens the file for reading and writing, and returns the resulting file.
// The filename is generated by taking pattern and adding a random string to the end.
// If pattern includes a "*", the random string replaces the last "*".
// If dir is the empty string, CreateTemp uses the default directory for temporary files, as returned by TempDir.
// Multiple programs or goroutines calling CreateTemp simultaneously will not choose the same file.
// The caller can use the file's Name method to find the pathname of the file.
// It is the caller's responsibility to remove the file when it is no longer needed.
func CreateTemp(dir, pattern string) (*File, error) {
	if dir == "" {
		dir = TempDir()
	}

	prefix, suffix, err := prefixAndSuffix(pattern)
	if err != nil {
		return nil, &PathError{Op: "createtemp", Path: pattern, Err: err}
	}
	prefix = joinPath(dir, prefix)

	attempt := 0
	for {
		name := prefix + nextRandom() + suffix
		f, err := OpenFile(name, O_RDWR|O_CREATE|O_EXCL, 0600)
		if IsExist(err) {
			if attempt++; attempt < 10000 {
				continue
			}
			return nil, &PathError{Op: "createtemp", Path: prefix + "*" + suffix, Err: ErrExist}
		}
		return f, err
	}
}

var errPatternHasSeparator = errors.New("pattern contains path separator")

// prefixAndSuffix splits pattern by the last wildcard "*", if applicable,
// returning prefix as the part before "*" and suffix as the part after "*".
func prefixAndSuffix(pattern string) (prefix, suffix string, err error) {
	for i := 0; i < len(pattern); i++ {
		if IsPathSeparator(pattern[i]) {
			return "", "", errPatternHasSeparator
		}
	}
	if pos := bytealg.LastIndexByteString(pattern, '*'); pos != -1 {
		prefix, suffix = pattern[:pos], pattern[pos+1:]
	} else {
		prefix = pattern
	}
	return prefix, suffix, nil
}

// MkdirTemp creates a new temporary directory in the directory dir
// and returns the pathname of the new directory.
// The new directory's name is generated by adding a random string to the end of pattern.
// If pattern includes a "*", the random string replaces the last "*" instead.
// If dir is the empty string, MkdirTemp uses the default directory for temporary files, as returned by TempDir.
// Multiple programs or goroutines calling MkdirTemp simultaneously will not choose the same directory.
// It is the caller's responsibility to remove the directory when it is no longer needed.
func MkdirTemp(dir, pattern string) (string, error) {
	if dir == "" {
		dir = TempDir()
	}

	prefix, suffix, err := prefixAndSuffix(pattern)
	if err != nil {
		return "", &PathError{Op: "mkdirtemp", Path: pattern, Err: err}
	}
	prefix = joinPath(dir, prefix)

	attempt := 0
	for {
		name := prefix + nextRandom() + suffix
		err := Mkdir(name, 0700)
		if err == nil {
			return name, nil
		}
		if IsExist(err) {
			if attempt++; attempt < 10000 {
				continue
			}
			return "", &PathError{Op: "mkdirtemp", Path: dir + string(PathSeparator) + prefix + "*" + suffix, Err: ErrExist}
		}
		if IsNotExist(err) {
			if _, err := Stat(dir); IsNotExist(err) {
				return "", err
			}
		}
		return "", err
	}
}

func joinPath(dir, name string) string {
	if len(dir) > 0 && IsPathSeparator(dir[len(dir)-1]) {
		return dir + name
	}
	return dir + string(PathSeparator) + name
}
